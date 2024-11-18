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
	HoeService *service.HoeService
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

	siteType := definitions.SiteType(*site)
	baseURL, ok := definitions.SiteConfigs[siteType]
	if !ok {
		return nil, fmt.Errorf("unknown site: %s", *site)
	}

	logger := InitLogger()
	baseConfig := definitions.ScraperConfig{
		Site:              siteType,
		BaseURL:           baseURL,
		RequestsPerSecond: 1.0,
		Logger:            logger,
	}

	scraper := scrapers.CreateScraper(baseConfig)

	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}
	hoeRepo := repository.NewHoeRepository(db)
	hoeService := service.NewHoeService(hoeRepo, logger)

	return &AppContext{
		Scraper:    scraper,
		Logger:     logger,
		HoeService: hoeService,
	}, nil
}
