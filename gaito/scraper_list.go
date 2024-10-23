package gaito

import (
	"context"
	"fmt"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

var BaseUrl = "https://www.gaito.mom"

func ProcessListPage(ctx context.Context) ([]string, error) {
	url := BaseUrl + "/gai-goi/khu-vuc/H%E1%BB%93%20Ch%C3%AD%20Minh/Qu%E1%BA%ADn%207"
	var urlList []string

	itemThreshold := 20
	evaluateScript := fmt.Sprintf(`new Promise((resolve, reject) => { const processItems = (items) => { return Array.from(items).map( (item) => item.querySelector(".thumbnail a").getAttribute("href") || "" ); }; const checkAndLoad = () => { const items = document.querySelectorAll( 'div[ng-repeat="item in products"]' ); if (items.length > %v) { resolve(processItems(items)); return; } const loadMoreButton = document.querySelector( "body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div:nth-child(4) > div > button" ); if (loadMoreButton) { loadMoreButton.click(); setTimeout(checkAndLoad, 2000); } else { resolve(processItems(items)); } }; checkAndLoad(); });`, itemThreshold)
	err := chromedp.Run(ctx, chromedp.Navigate(url),
		chromedp.WaitVisible(`div[ng-repeat="item in products"]`, chromedp.ByQueryAll),
		chromedp.Evaluate(evaluateScript, &urlList, func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
			return p.WithAwaitPromise(true)
		}),
	)

	if err != nil {
		return nil, err
	}

	return urlList, nil
}
