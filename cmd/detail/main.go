package main

import (
	"github.com/HumbeBee/hoe-crawler/internal/utils/setuputil"
)

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	// Example url
	url := "/gai-goi/chi-tietz/56042/hot-girl-diep-anhmat-xinh-nguc-dep-bu-cu-dieu-luyen"

	err = appContext.HoeService.ProcessDetailPage(url)
	if err != nil {
		appContext.FailedUrlService.TrackFailedURL(url, err)
	}
}
