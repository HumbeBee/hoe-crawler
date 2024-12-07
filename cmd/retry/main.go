package retry

import "github.com/HumbeBee/hoe-crawler/internal/utils/setuputil"

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	err = appContext.FailedUrlService.RetryFailedURLs()
	if err != nil {
		appContext.Logger.Error(err.Error())
	}
}
