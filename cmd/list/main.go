package main

import (
	"log"

	"github.com/haovoanh28/gai-webscraper/cmd/config"
	"github.com/haovoanh28/gai-webscraper/internal/sites/gaito"
)

func main() {
	config.InitLogger()

	scraper := gaito.NewScraper(config.BaseUrl, config.RequestPerSecond, log.Default())
	scraper.ProcessListPage()
}
