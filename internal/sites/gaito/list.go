package gaito

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func (s *Scraper) processListPage() ([]string, error) {
	const (
		itemThreshold = 30
	)

	url := s.baseURL + "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"
	var urlList []string
	page := rod.New().Timeout(30 * time.Second).MustConnect().MustPage(url).MustWaitStable()
	defer page.Close()

	fmt.Println("Loading...", page.MustInfo())

	for {
		items := page.MustElements(listPageSelectors.Items)
		currentLength := len(items)

		// currentLength >= itemThreshold: enough items
		// currentLength == 0: for some reason, the query doesn't return any items (Ex: Cloudflare, ...)
		if currentLength >= itemThreshold || currentLength == 0 {
			break
		}

		loadMoreBtn := page.MustElement(listPageSelectors.LoadMoreBtn)
		loadMoreBtn.MustClick()

		page.MustWaitElementsMoreThan(listPageSelectors.Items, currentLength)
	}

	elements := page.MustElements(listPageSelectors.Items)
	fmt.Println("Found", len(elements), "items")
	for _, elem := range elements {
		urlList = append(urlList, *elem.MustElement(listPageSelectors.ThumbnailUrl).MustAttribute("href"))
	}

	return urlList, nil
}
