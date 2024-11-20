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
	Logger     *logutil.Logger
	HoeService service.HoeService
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
	site := flag.String("site", "", "The site to scrape")
	flag.Parse()

	logger := InitLogger()

	db, err := database.InitDB()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	siteRepo := repository.NewSiteRepository(db)
	siteInfo, err := siteRepo.GetSiteByName(*site)
	if err != nil {
		return nil, fmt.Errorf("failed to get site by name: %v", err)
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
	hoeService := service.NewHoeService(hoeRepo, logger, scraper)

	return &AppContext{
		Scraper:    scraper,
		Logger:     logger,
		HoeService: hoeService,
	}, nil
}
