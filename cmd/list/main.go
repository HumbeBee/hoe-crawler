package main

import (
	"log"

	"github.com/haovoanh28/gai-webscraper/cmd/config"
	"github.com/haovoanh28/gai-webscraper/internal/sites/gaito"
)

func main() {
	config.InitLogger()

	log.Println("Starting URL list scraper ...")

	scraper := gaito.NewScraper(config.BaseUrl, config.RequestPerSecond)

	urls := scraper.GetList()
}
