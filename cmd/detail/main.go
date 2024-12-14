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
	url := "/gai-goi/chi-tiet/63863/reup-hot-girl-huong-ly-mat-xinh-goi-tinh-body-dep-nhin-rat-sexy"

	err = appContext.HoeService.ProcessDetailPage(appContext.SiteInfo.BaseURL, url)
	if err != nil {
		appContext.FailedUrlService.TrackFailedURL(models.FailedTypeDetail, url, err)
	}
}
