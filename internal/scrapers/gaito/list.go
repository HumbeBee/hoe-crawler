package gaito

import (
	"fmt"
	"github.com/go-rod/rod/lib/proto"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
	"time"
)

func (s *Scraper) ProcessListPage() ([]string, error) {
	const (
		itemThreshold = 30
	)
	url := s.BaseURL + "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"

	s.Logger.Info(fmt.Sprintf("Processing %s", url))

	rodBrowser, page, root, err := browser.ConnectToPage(url, 30*time.Second)
	if err != nil {
		return nil, errutil.WrapError("connect to page", err, url)
	}
	defer rodBrowser.Close()
	defer page.Close()

	var urlList []string
	for {
		items, err := browser.GetMultipleElementsWithRetry(root, listPageSelectors.Items)
		if err != nil {
			return nil, errutil.WrapError("get list items", err, url)
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
			return nil, errutil.WrapError("click load more button", err, url)
		}

		if err := page.WaitElementsMoreThan(listPageSelectors.Items, currentLength); err != nil {
			return nil, errutil.WrapError("wait more items", err, url)
		}
	}

	elements, err := browser.GetMultipleElementsWithRetry(root, listPageSelectors.Items)
	if err != nil {
		return nil, errutil.WrapError("get final list items", err, url)
	}

	for _, elem := range elements {
		urlList = append(urlList, *elem.MustElement(listPageSelectors.ThumbnailUrl).MustAttribute("href"))
	}

	return urlList, nil
}
