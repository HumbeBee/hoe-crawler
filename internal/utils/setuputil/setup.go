package setuputil

import (
	"flag"
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/config"
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/database"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
	"github.com/HumbeBee/hoe-crawler/internal/scrapers"
	"github.com/HumbeBee/hoe-crawler/internal/service"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
	"log"
	"time"
)

type AppContext struct {
	Logger           *logutil.Logger
	Scraper          interfaces.Scraper
	HoeService       interfaces.HoeService
	FailedUrlService interfaces.FailedURLService
}

func InitLogger() *logutil.Logger {
	log.SetFlags(log.LstdFlags)

	envConfig := config.GetEnvConfig()
	logLevelStr := envConfig.LOGLEVEL
	logLevel, err := logutil.ParseLogLevel(logLevelStr)
	if err != nil {
		log.Printf("Invalid log level '%s', defaulting to INFO", logLevelStr)
		logLevel = logutil.INFO
	}

	return logutil.NewLogger(logLevel)
}

func CreateAppContext() (*AppContext, error) {
	// Get site from cmd options
	site := flag.String("site", "gaito", "The site to scrape")
	flag.Parse()

	logger := InitLogger()

	db, err := database.InitDB()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	siteRepo := repository.NewSiteRepository(db)
	siteInfo, err := siteRepo.GetSiteByName(*site)
	if err != nil {
		return nil, fmt.Errorf("failed to get site by name: %w", err)
	}

	baseScraperConfig := definitions.ScraperConfig{
		SiteID:            siteInfo.ID,
		SiteName:          siteInfo.Name,
		BaseURL:           siteInfo.BaseURL,
		RequestsPerSecond: 1.0,
		Logger:            logger,
	}

	scraper := scrapers.CreateScraper(baseScraperConfig)

	hoeRepo := repository.NewHoeRepository(db, logger)
	locationRepo := repository.NewLocationRepository(db)
	workingHistoryRepo := repository.NewWorkingHistoryRepository(db, logger)

	browserRateLimiter := browser.NewBrowserRateLimiter(3 * time.Second)

	hoeService, err := service.NewHoeBuilder().
		WithHoeRepo(hoeRepo).
		WithWorkingHistoryRepo(workingHistoryRepo).
		WithLocationRepo(locationRepo).
		WithLogger(logger).
		WithScraper(scraper).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create hoe failedURLService: %w", err)
	}

	failedURLRepo := repository.NewFailedURLRepository(db)
	failedURLService, err := service.NewFailedURLBuilder().
		WithSiteID(siteInfo.ID).
		WithLogger(logger).
		WithBrowserRateLimiter(browserRateLimiter).
		WithFailedURLRepo(failedURLRepo).
		WithSiteRepo(siteRepo).
		WithHoeService(hoeService).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create failedURLService: %w", err)
	}

	return &AppContext{
		Scraper:          scraper,
		Logger:           logger,
		HoeService:       hoeService,
		FailedUrlService: failedURLService,
	}, nil
}
