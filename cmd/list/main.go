package main

import (
	"fmt"
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
		logger.Info(fmt.Sprintf("Found %d items\n", len(urlList)))

		for _, url := range urlList {
			logger.Info(url)
		}
	}
}
