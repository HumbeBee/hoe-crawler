package main

import (
	"github.com/haovoanh28/gai-webscraper/internal/utils/setuputil"
)

func main() {
	scraper, logger := setuputil.SetupEnvironment()

	// Example url
	url := "/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	hoe, err := scraper.ProcessDetailPage(url)
	if err != nil {
		logger.Fatal(err.Error())
	}

	if hoe != nil {
		hoe.Print()
	}
}
