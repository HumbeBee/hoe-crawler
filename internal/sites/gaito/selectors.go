package gaito

// ListPageSelectors contains all selectors used in list page scraping
type ListPageSelectors struct {
	Items        string
	LoadMoreBtn  string
	ThumbnailUrl string
}

var (
	listPageSelectors = ListPageSelectors{
		Items:        `div[ng-repeat="item in products"]`,
		LoadMoreBtn:  `body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div:nth-child(4) > div > button`,
		ThumbnailUrl: `.thumbnail a`,
	}
)
