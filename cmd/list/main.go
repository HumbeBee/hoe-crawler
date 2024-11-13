package main

import (
	"github.com/haovoanh28/gai-webscraper/internal/utils/setuputil"
)

func main() {
	scraper, logger := setuputil.SetupEnvironment()

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
