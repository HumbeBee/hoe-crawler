package main

import (
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"github.com/HumbeBee/hoe-crawler/internal/utils/setuputil"
)

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	// Example url
	url := "/gai-goi/chi-tiet/63459/baby-tram-anh-tre-trung-xinh-xan-lan-dau-len-song"

	err = appContext.HoeService.ProcessDetailPage(appContext.SiteInfo.BaseURL, url)
	if err != nil {
		appContext.FailedUrlService.TrackFailedURL(models.FailedTypeDetail, url, err)
	}
}
