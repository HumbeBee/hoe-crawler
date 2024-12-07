package gaito

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
)

type listPageScraper struct {
	conn *browser.Connection
	url  string
}

func newListPageScraper(conn *browser.Connection, url string) *listPageScraper {
	return &listPageScraper{conn: conn, url: url}
}

func (s *listPageScraper) getHoeURLs() ([]string, error) {
	const (
		itemThreshold = 30
	)
	var urlList []string
	for {
		items, err := s.conn.Root.FindAll(listPageSelectors.Items)
		if err != nil {
			return nil, fmt.Errorf("get list items: %w", err)
		}

		// currentLength >= itemThreshold: enough items
		// currentLength == 0: for some reason, the query doesn't return any items (Ex: Cloudflare, ...)
		currentLength := len(items)
		if currentLength >= itemThreshold || currentLength == 0 {
			break
		}

		loadMoreBtn, err := s.conn.Root.Find(listPageSelectors.LoadMoreBtn)
		if err != nil {
			break
		}

		if err := loadMoreBtn.Click(); err != nil {
			return nil, fmt.Errorf("click load more button: %w", err)
		}

		if err := s.conn.Page.WaitElementsMoreThan(listPageSelectors.Items, currentLength); err != nil {
			return nil, fmt.Errorf("wait more items: %w", err)
		}
	}

	elements, err := s.conn.Root.FindAll(listPageSelectors.Items)
	if err != nil {
		return nil, fmt.Errorf("get final list items: %w", err)
	}

	for _, elem := range elements {
		urlList = append(urlList, elem.MustFind(listPageSelectors.ThumbnailUrl).MustGetAttribute("href"))
	}

	return urlList, nil
}
