package gaito

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
)

func (s *Scraper) processListPage() ([]string, error) {
	const (
		itemThreshold = 30
		stableTimeout = 30
	)
	url := s.baseURL + "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"

	fmt.Println("Starting to load", url)

	rodBrowser := rod.New()
	if err := rodBrowser.Connect(); err != nil {
		return nil, s.ErrorHandler.WrapError("connect browser", err, url)
	}
	defer rodBrowser.Close()

	page, err := rodBrowser.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		return nil, s.ErrorHandler.WrapError("create page", err, url)
	}
	defer page.Close()

	if err := page.WaitStable(time.Duration(stableTimeout)); err != nil {
		return nil, s.ErrorHandler.WrapError("wait stable", err, url)
	}

	root := page.MustElement("html")
	var urlList []string
	for {
		items, err := browser.GetMultipleElementsWithRetry(root, listPageSelectors.Items)
		if err != nil {
			return nil, s.ErrorHandler.WrapError("get list items", err, url)
		}

		// currentLength >= itemThreshold: enough items
		// currentLength == 0: for some reason, the query doesn't return any items (Ex: Cloudflare, ...)
		currentLength := len(items)
		if currentLength >= itemThreshold || currentLength == 0 {
			break
		}

		loadMoreBtn, err := browser.GetElementWithRetry(root, listPageSelectors.LoadMoreBtn)
		if err != nil {
			break
		}

		if err := loadMoreBtn.Click(proto.InputMouseButtonLeft, 1); err != nil {
			return nil, s.ErrorHandler.WrapError("click load more", err, url)
		}

		if err := page.WaitElementsMoreThan(listPageSelectors.Items, currentLength); err != nil {
			return nil, s.ErrorHandler.WrapError("wait for more items", err, url)
		}
	}

	elements, err := browser.GetMultipleElementsWithRetry(root, listPageSelectors.Items)
	if err != nil {
		return nil, s.ErrorHandler.WrapError("get final list items", err, url)
	}

	for _, elem := range elements {
		urlList = append(urlList, *elem.MustElement(listPageSelectors.ThumbnailUrl).MustAttribute("href"))
	}

	return urlList, nil
}
