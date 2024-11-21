package gaito

import (
	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/browser"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
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
			return nil, errutil.WrapError("get list items", err, s.url)
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
			return nil, errutil.WrapError("click load more button", err, s.url)
		}

		if err := s.conn.Page.WaitElementsMoreThan(listPageSelectors.Items, currentLength); err != nil {
			return nil, errutil.WrapError("wait more items", err, s.url)
		}
	}

	elements, err := s.conn.Root.FindAll(listPageSelectors.Items)
	if err != nil {
		return nil, errutil.WrapError("get final list items", err, s.url)
	}

	for _, elem := range elements {
		urlList = append(urlList, elem.MustFind(listPageSelectors.ThumbnailUrl).MustGetAttribute("href"))
	}

	return urlList, nil
}
