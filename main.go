package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
	"github.com/haovoanh28/gai-webscraper/gaito"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// itemThreshold := 100
	urlList, err := gaito.ProcessListPage(ctx)

	if err != nil {
		panic(err)
	}

	for _, url := range urlList {
		// gaito.ProcessDetailPage(ctx, url)
		fmt.Println(url)
	}

	// From now, just process the first page to make it works first
	hoe, err := gaito.ProcessDetailPage(ctx, urlList[0])

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", hoe)
}
