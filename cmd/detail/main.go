package main

import (
	"github.com/haovoanh28/gai-webscraper/internal/scrapers"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

func main() {
	logger := logutil.InitLogger()

	// Example url
	url := "/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	scraper, err := scrapers.CreateScraper("gaito")
	if err != nil {
		logger.Fatal(err.Error())
	}

	hoe, err := scraper.ProcessDetailPage(url)
	if err != nil {
		logger.Fatal(err.Error())
	}
	hoe.Print()
}
