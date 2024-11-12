package main

import (
	"fmt"
	"log"

	"github.com/haovoanh28/gai-webscraper/cmd/config"
	"github.com/haovoanh28/gai-webscraper/internal/scrapers"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
)

func main() {
	config.InitLogger()

	errorHandler := errutil.NewErrorHandler(log.Default(), errutil.DEBUG)
	scraper, err := scrapers.CreateScraper("gaito")
	if err != nil {
		panic(err)
	}
	urlList, err := scraper.ProcessListPage()
	if err != nil {
		errorHandler.Fatal(err.Error())
	}

	if len(urlList) == 0 {
		fmt.Println("No items found (Maybe Cloudflare block ?)")
	} else {
		log.Println("Found", len(urlList), "items")
	}
}
