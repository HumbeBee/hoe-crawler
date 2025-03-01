package gaito

// ListPageSelectors contains all selectors used in list page scraping
type ListPageSelectors struct {
	Items        string
	LoadMoreBtn  string
	ThumbnailUrl string
}

type DetailPageSelectors struct {
	PageContainer  string
	DetailInfoTab  string
	DetailInfoCard string

	ReportTab           string
	ReportTabContent    string
	ReportList          string
	ReportViewMoreBtn   string
	ReportGoNextPageBtn string

	CityName     string
	DistrictName string
	Name         string
	ImageUrl     string
	Price        string
	Phone        string
	Address      string
	Author       string
	Status       string
	BirthYear    string
	Height       string
	Weight       string
	Country      string
	Service      string
	Duration     string
	WorkTime     string
}

type ReportPageSelectors struct {
}

var (
	listPageSelectors = ListPageSelectors{
		Items:        `div[ng-repeat="item in products"]`,
		LoadMoreBtn:  `body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div:nth-child(4) > div > button`,
		ThumbnailUrl: `.thumbnail a`,
	}

	detailPageSelectors = DetailPageSelectors{
		PageContainer:  `.container.seduction-container .ow_page_container`,
		DetailInfoTab:  `.tab-content`,
		DetailInfoCard: `.jumbotron.ng-scope`,

		ReportTab:           `li[index="2"] a.nav-link`,
		ReportTabContent:    `product-review[ng-if="reviewTabLoaded"] div.panel-body`,
		ReportList:          `div[ng-repeat="review in reviews"]`,
		ReportViewMoreBtn:   `a.view_more_report`,
		ReportGoNextPageBtn: `product-review li.pagination-next:not(.disabled) a[ng-click]`,

		CityName:     `.breadcrumb > li:nth-child(3) > a > span`,
		DistrictName: `.breadcrumb > li:nth-child(4) > a > span`,
		Name:         `div:nth-child(3) > div > h1`,
		ImageUrl:     `div:nth-child(3) > div > div:nth-child(3) > div > div > div > div.tab-pane.ng-scope.active > div.jumbotron.ng-scope > div > div.col-md-3.col-sm-4.media.escort_item_wrap > div > image-placeholder > img`,
		Price:        `.jumbotron .fa.fa-money + span`,
		Phone:        `.jumbotron .fa.fa-phone + a`,
		Address:      `.jumbotron .fa.fa-map-marker + a`,
		Author:       `.jumbotron .fa.fa-user + a`,
		Status:       `.jumbotron .fa.fa-file-o + span`,
		BirthYear:    `product-attribute table > tbody > tr:nth-child(3) > td:nth-child(2) > attribute-dob-box > div > div`,
		Height:       `product-attribute table > tbody > tr:nth-child(4) > td:nth-child(2) > attribute-number-box .ng-scope`,
		Weight:       `product-attribute table > tbody > tr:nth-child(5) > td:nth-child(2) > attribute-number-box .ng-scope`,
		Country:      `product-attribute table > tbody > tr:nth-child(9) > td:nth-child(2) > attribute-radio-box span span[ng-repeat="item in attributeDto.settings.values"]`,
		Service:      `product-attribute table > tbody > tr:nth-child(12) > td:nth-child(2) > attribute-choices-box span span`,
		Duration:     `product-attribute table > tbody > tr:nth-last-child(3) > td:nth-child(2) > attribute-text-box span`,
		WorkTime:     `product-attribute table > tbody > tr:nth-last-child(2) > td:nth-child(2) > attribute-text-box span`,
	}
)
