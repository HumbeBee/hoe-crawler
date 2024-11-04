package gaito

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

var BaseUrl = "https://www.gaito.love"

func ProcessListPage() []string {
	url := BaseUrl + "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"
	var urlList []string

	itemThreshold := 30

	page := rod.New().Timeout(30 * time.Second).MustConnect().MustPage(url).MustWaitStable()
	defer page.Close()
	itemsQuery := `div[ng-repeat="item in products"]`

	fmt.Println("Loading...", page.MustInfo())

	for {
		items := page.MustElements(itemsQuery)
		currentLength := len(items)

		// currentLength >= itemThreshold: enough items
		// currentLength == 0: for some reason, the query doesn't return any items (Ex: Cloudflare, ...)
		if currentLength >= itemThreshold || currentLength == 0 {
			break
		}

		loadMoreBtn := page.MustElement(`body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div:nth-child(4) > div > button`)
		loadMoreBtn.MustClick()

		page.MustWaitElementsMoreThan(itemsQuery, currentLength)
	}

	elements := page.MustElements(itemsQuery)
	fmt.Println("Found", len(elements), "items")
	for _, elem := range elements {
		urlList = append(urlList, *elem.MustElement(".thumbnail a").MustAttribute("href"))
	}

	return urlList
}
