package gaito

import (
	"fmt"
	"strings"
	"time"

	"github.com/haovoanh28/gai-webscraper/internal/hoe"
	"github.com/haovoanh28/gai-webscraper/utils"

	"github.com/go-rod/rod"
)

func ProcessDetailUrl(url string) *hoe.Hoe {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error processing detail URL %s: %v\n", url, err)
		}
	}()

	result := ProcessDetailPage(url)
	return &result
}

func ProcessDetailPage(detailUrl string) hoe.Hoe {
	url := BaseUrl + detailUrl

	id := utils.GetIDFromUrl(detailUrl)

	// Wait until content element is visible
	page := rod.New().MustConnect().MustPage(url)
	containerElement := page.MustElement(`.container.seduction-container .ow_page_container`).MustWaitVisible()
	detailInfoElement := page.MustElement(`.tab-content`).MustWaitVisible()

	var mainInfo hoe.HoeMainInfo
	mainInfo.Name = utils.GetElementsText(containerElement, `body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div > h1`, "name")
	mainInfo.ImageUrl = utils.GetElementAttribute(containerElement, `body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div > div:nth-child(3) > div > div > div > div.tab-pane.ng-scope.active > div.jumbotron.ng-scope > div > div.col-md-3.col-sm-4.media.escort_item_wrap > div > image-placeholder > img`, "src", "image_url")
	// Ex: "300 k" => "300k"
	price := utils.GetElementText(detailInfoElement, `.jumbotron .fa.fa-money + span`, "price")
	mainInfo.Price = strings.ReplaceAll(price, "\u00A0", "")
	// Ex: "0123.456.789" -> "0123456789"
	phone := utils.GetElementText(detailInfoElement, `.jumbotron .fa.fa-phone + a`, "phone")
	mainInfo.Phone = strings.ReplaceAll(phone, ".", "")
	mainInfo.Address = utils.GetElementText(detailInfoElement, `.jumbotron .fa.fa-map-marker + a`, "address")
	mainInfo.Author = utils.GetElementText(detailInfoElement, `.jumbotron .fa.fa-user + a`, "author")
	mainInfo.Status = utils.GetElementText(detailInfoElement, `.jumbotron .fa.fa-file-o + span`, "status")

	var detailInfo hoe.HoeDetailInfo
	detailInfoElement = page.MustElement(`product-attribute .table-responsive`).MustWaitVisible()
	time.Sleep(2 * time.Second)
	detailInfo.BirthYear = utils.GetElementText(detailInfoElement, `body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div > div:nth-child(3) > div > div > div > div.tab-pane.ng-scope.active > div:nth-child(2) > product-attribute > div > div > div > table > tbody > tr:nth-child(3) > td:nth-child(2) > attribute-dob-box > div > div`, "birth_year")
	detailInfo.Height = utils.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(4) > td:nth-child(2) > attribute-number-box .ng-scope`, "height") + "cm"
	detailInfo.Weight = utils.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(5) > td:nth-child(2) > attribute-number-box .ng-scope`, "weight") + "kg"
	detailInfo.From = utils.GetElementsText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(9) > td:nth-child(2) > attribute-radio-box span span[ng-repeat="item in attributeDto.settings.values"]`, "from")
	detailInfo.Service = utils.GetElementsText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(12) > td:nth-child(2) > attribute-choices-box span span[ng-repeat="item in attributeDto.settings.values"]`, "service")
	detailInfo.Duration = utils.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(15) > td:nth-child(2) > attribute-text-box span`, "duration")
	detailInfo.WorkTime = utils.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(16) > td:nth-child(2) > attribute-text-box span`, "work_time")

	// Get report urls
	var reportUrls []string
	page.MustElement(`li[index="2"] a.nav-link`).MustClick().MustWaitLoad()
	time.Sleep(1 * time.Second)
	reportTabElement := page.MustElement(`product-review[ng-if="reviewTabLoaded"]`)

	for {
		reportElements, err := reportTabElement.Elements(`div[ng-repeat="review in reviews"]`)
		if err != nil {
			panic(err)
		}
		if len(reportElements) == 0 && len(reportUrls) == 0 {
			continue
			// panic(fmt.Errorf("empty reportElements ?: %v", err))
		}

		for _, reportElement := range reportElements {
			btnElement, err := reportElement.Element(`a.view_more_report`)
			if err != nil {
				panic(fmt.Errorf("failed to get view_more_report: %v", err))
			}

			reportUrl, err := btnElement.Attribute("href")
			if err != nil {
				panic(fmt.Errorf("failed to get reportUrl: %v", err))
			}
			reportUrls = append(reportUrls, *reportUrl)
		}

		if len(reportUrls) == 0 {
			break
		}

		goNextPageBtn, err := page.Timeout(1 * time.Second).Element(`product-review li.pagination-next:not(.disabled) a[ng-click]`)
		if err != nil {
			break
		} else {
			goNextPageBtn.MustClick().MustWaitLoad().CancelTimeout()
			time.Sleep(1 * time.Second)
		}
	}

	return hoe.Hoe{ID: id, Url: url, MainInfo: &mainInfo, DetailInfo: &detailInfo, ReportURLs: reportUrls}
}
