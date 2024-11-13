package main

import (
	"github.com/haovoanh28/gai-webscraper/internal/scrapers"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

func main() {
	logger := logutil.InitLogger()

	scraper, err := scrapers.CreateScraper("gaito", logger)
	if err != nil {
		logger.Fatal("create scraper", err.Error())
	}

	urlList, err := scraper.ProcessListPage()
	if err != nil {
		logger.Fatal("process list page", err.Error())
	}

	if len(urlList) == 0 {
		logger.Warn("No items found (Maybe Cloudflare block ?)")
	} else {
		logger.Info("Found", len(urlList), "items")
	}
}
