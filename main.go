package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	url := "https://www.gaito.mom/gai-goi/khu-vuc/H%E1%BB%93%20Ch%C3%AD%20Minh/Qu%E1%BA%ADn%207"
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlContents []string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`div[ng-repeat="item in products"]`, chromedp.ByQueryAll),
		// chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('div[ng-repeat="item in products"]')).map(e => e.outerHTML)`, &htmlContents),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, content := range htmlContents {
		fmt.Println(content)
	}
}
