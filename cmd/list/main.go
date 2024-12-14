package main

import (
	"github.com/HumbeBee/hoe-crawler/internal/utils/setuputil"
)

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	// Example URL
	url := "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"

	err = appContext.HoeService.ProcessListPage(appContext.SiteInfo.BaseURL, url)
	if err != nil {
		appContext.Logger.Fatal(err.Error())
	}
}
