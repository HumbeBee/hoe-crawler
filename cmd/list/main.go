package main

import (
	"github.com/haovoanh28/gai-webscraper/internal/utils/setuputil"
)

func main() {
	appContext, err := setuputil.CreateAppContext()
	if err != nil {
		panic(err)
	}

	appContext.HoeService.ProcessListPage()
}
