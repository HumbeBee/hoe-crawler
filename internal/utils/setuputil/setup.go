package setuputil

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/database"
	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
	"github.com/haovoanh28/gai-webscraper/internal/scrapers"
	"github.com/haovoanh28/gai-webscraper/internal/service"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

type AppContext struct {
	Scraper    interfaces.Scraper
	HoeService service.HoeService
	Logger     *logutil.Logger
}

func InitLogger() *logutil.Logger {
	log.SetFlags(log.LstdFlags)

	logLevelStr := os.Getenv("LOG_LEVEL")
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

	baseConfig := definitions.ScraperConfig{
		SiteID:            siteInfo.ID,
		SiteName:          siteInfo.Name,
		BaseURL:           siteInfo.BaseURL,
		RequestsPerSecond: 1.0,
		Logger:            logger,
	}

	scraper := scrapers.CreateScraper(baseConfig)

	hoeRepo := repository.NewHoeRepository(db, logger)
	locationRepo := repository.NewLocationRepository(db)
	workingHistoryRepo := repository.NewWorkingHistoryRepository(db, logger)

	hoeService, err := service.NewHoeBuilder().
		WithHoeRepo(hoeRepo).
		WithWorkingHistoryRepo(workingHistoryRepo).
		WithLocationRepo(locationRepo).
		WithLogger(logger).
		WithScraper(scraper).
		Build()

	if err != nil {
		return nil, fmt.Errorf("failed to create hoe service: %w", err)
	}

	return &AppContext{
		Scraper:    scraper,
		Logger:     logger,
		HoeService: hoeService,
	}, nil
}
