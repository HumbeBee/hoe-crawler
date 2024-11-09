package main

import (
	"fmt"
	"log"

	"github.com/haovoanh28/gai-webscraper/cmd/config"
	"github.com/haovoanh28/gai-webscraper/internal/sites/gaito"
)

func main() {
	config.InitLogger()

	scraper := gaito.NewScraper(config.BaseUrl, config.RequestPerSecond, log.Default())
	urls, err := scraper.GetList()
	if err != nil {
		scraper.ErrorHandler.Fatal(err.Error())
	}

	if len(urls) == 0 {
		scraper.ErrorHandler.Info("No items found (Maybe Cloudflare block ?)")
	} else {
		for _, url := range urls {
			fmt.Println(url)
		}
	}
}
