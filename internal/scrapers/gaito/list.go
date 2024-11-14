package gaito

import (
	"fmt"
	"time"

	"github.com/go-rod/rod/lib/proto"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
)

func (s *scraper) ProcessListPage() ([]string, error) {
	const (
		itemThreshold = 30
	)
	url := s.BaseURL + "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"

	s.Logger.Info(fmt.Sprintf("Processing %s", url))

	conn, err := browser.ConnectToPage(url, 30*time.Second)
	if err != nil {
		return nil, errutil.WrapError("connect to page", err, url)
	}
	defer conn.Close()

	var urlList []string
	for {
		items, err := browser.GetMultipleElementsWithRetry(conn.Root, listPageSelectors.Items)
		if err != nil {
			return nil, errutil.WrapError("get list items", err, url)
		}

		// currentLength >= itemThreshold: enough items
		// currentLength == 0: for some reason, the query doesn't return any items (Ex: Cloudflare, ...)
		currentLength := len(items)
		if currentLength >= itemThreshold || currentLength == 0 {
			break
		}

		loadMoreBtn, err := browser.GetElementWithRetry(conn.Root, listPageSelectors.LoadMoreBtn)
		if err != nil {
			break
		}

		if err := loadMoreBtn.Click(proto.InputMouseButtonLeft, 1); err != nil {
			return nil, errutil.WrapError("click load more button", err, url)
		}

		if err := conn.Page.WaitElementsMoreThan(listPageSelectors.Items, currentLength); err != nil {
			return nil, errutil.WrapError("wait more items", err, url)
		}
	}

	elements, err := browser.GetMultipleElementsWithRetry(conn.Root, listPageSelectors.Items)
	if err != nil {
		return nil, errutil.WrapError("get final list items", err, url)
	}

	for _, elem := range elements {
		urlList = append(urlList, *elem.MustElement(listPageSelectors.ThumbnailUrl).MustAttribute("href"))
	}

	return urlList, nil
}
