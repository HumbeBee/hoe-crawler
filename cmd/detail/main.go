package main

import (
	"github.com/haovoanh28/gai-webscraper/internal/utils/setuputil"
)

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	// Example url
	url := "/gai-goi/chi-tiet/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	err = appContext.HoeService.ProcessDetailPage(url)
	if err != nil {
		appContext.Logger.Fatal(err.Error())
	}
}
