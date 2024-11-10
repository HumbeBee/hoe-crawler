package main

import (
	"log"

	"github.com/haovoanh28/gai-webscraper/cmd/config"
	"github.com/haovoanh28/gai-webscraper/internal/sites/gaito"
)

func main() {
	config.InitLogger()

	// Example url
	// url := "https://www.gaito.love/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"
	url := "/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	scraper := gaito.NewScraper(config.BaseUrl, config.RequestPerSecond, log.Default())

	hoeInfo, err := scraper.GetDetail(url)
	if err != nil {
		scraper.ErrorHandler.Fatal(err.Error())
	}

	// Save hoeInfo to db

	if len(hoeInfo.ReportUrls) > 0 {
		// Should save urls into db
	}

	hoeInfo.Print()
}
