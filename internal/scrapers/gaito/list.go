package gaito

import (
	"fmt"

	"github.com/go-rod/rod/lib/proto"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
)

func (s *gaitoScraper) ProcessListPage() ([]string, error) {
	const (
		itemThreshold = 30
	)
	url := s.BaseURL + "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"

	fmt.Println("Starting to load", url)

	rodBrowser, page, root, err := browser.ConnectToPage(url)
	if err != nil {
		// return nil, s.ErrorHandler.WrapError("connect to page", err, url)
		return nil, fmt.Errorf("connect to page %s: %w", url, err)
	}
	defer rodBrowser.Close()
	defer page.Close()

	var urlList []string
	for {
		items, err := browser.GetMultipleElementsWithRetry(root, listPageSelectors.Items)
		if err != nil {
			return nil, fmt.Errorf("get list items %s: %w", url, err)
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
			return nil, fmt.Errorf("click load more button %s: %w", url, err)
		}

		if err := page.WaitElementsMoreThan(listPageSelectors.Items, currentLength); err != nil {
			return nil, fmt.Errorf("wait more items %s: %w", url, err)
		}
	}

	elements, err := browser.GetMultipleElementsWithRetry(root, listPageSelectors.Items)
	if err != nil {
		return nil, fmt.Errorf("get final list items %s: %w", url, err)
	}

	for _, elem := range elements {
		urlList = append(urlList, *elem.MustElement(listPageSelectors.ThumbnailUrl).MustAttribute("href"))
	}

	return urlList, nil
}
