package main

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/utils/setuputil"
)

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	urlList, err := appContext.Scraper.ProcessListPage()
	if err != nil {
		appContext.Logger.Fatal("process list page", err.Error())
	}

	if len(urlList) == 0 {
		appContext.Logger.Warn("No items found (Maybe Cloudflare block ?)")
	} else {
		appContext.Logger.Info(fmt.Sprintf("Found %d items\n", len(urlList)))

		for _, url := range urlList {
			appContext.Logger.Info(url)
		}
	}
}
