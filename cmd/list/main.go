package main

import (
	"github.com/HumbeBee/hoe-crawler/internal/utils/setuputil"
)

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	err = appContext.HoeService.ProcessListPage()
	if err != nil {
		appContext.Logger.Fatal(err.Error())
	}
}
