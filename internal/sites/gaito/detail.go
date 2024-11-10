package gaito

import (
	"strings"

	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/utils"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
)

func (s *Scraper) processDetailPage(detailUrl string) (*models.HoeInfo, error) {
	url := s.baseURL + detailUrl

	id := utils.GetIDFromUrl(detailUrl)

	// Wait until content element is visible
	rodBrowser, page, root, err := browser.ConnectToPage(url)
	if err != nil {
		return nil, s.ErrorHandler.WrapError("connect to page", err, url)
	}
	defer rodBrowser.Close()
	defer page.Close()

	containerEle, err := browser.GetVisibleElement(root, detailPageSelectors.PageContainer)
	if err != nil {
		return nil, s.ErrorHandler.WrapError("get container element", err, url)
	}

	detailInfoTabEle, err := browser.GetVisibleElement(containerEle, detailPageSelectors.DetailInfoTab)
	if err != nil {
		return nil, s.ErrorHandler.WrapError("get detail info tab element", err, url)
	}

	hoeInfo := models.HoeInfo{
		Url: url,
		ID:  id,
	}

	hoeInfo.Name = browser.MustGetElementText(containerEle, detailPageSelectors.Name)

	hoeInfo.ImageUrl = browser.MustGetElementAttribute(containerEle, detailPageSelectors.ImageUrl, "src")

	// Ex: "300 k" => "300k"
	price := browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Price)
	hoeInfo.Price = strings.ReplaceAll(price, "\u00A0", "")

	// Ex: "0123.456.789" -> "0123456789"
	phone := browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Phone)
	hoeInfo.Phone = strings.ReplaceAll(phone, ".", "")

	hoeInfo.Address = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Address)
	hoeInfo.Provider = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Author)
	hoeInfo.Status = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Status)

	// detailInfoElement = page.MustElement(`product-attribute .table-responsive`).MustWaitVisible()
	// time.Sleep(2 * time.Second)
	// hoeInfo.BirthYear = browser.GetElementText(detailInfoElement, `body > div.container.seduction-container > div.knn_page_wrap > div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div > div:nth-child(3) > div > div > div > div.tab-pane.ng-scope.active > div:nth-child(2) > product-attribute > div > div > div > table > tbody > tr:nth-child(3) > td:nth-child(2) > attribute-dob-box > div > div`, "birth_year")
	// hoeInfo.Height = browser.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(4) > td:nth-child(2) > attribute-number-box .ng-scope`, "height") + "cm"
	// hoeInfo.Weight = browser.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(5) > td:nth-child(2) > attribute-number-box .ng-scope`, "weight") + "kg"
	// hoeInfo.From = browser.GetElementsText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(9) > td:nth-child(2) > attribute-radio-box span span[ng-repeat="item in attributeDto.settings.values"]`, "from")
	// hoeInfo.Service = browser.GetElementsText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(12) > td:nth-child(2) > attribute-choices-box span span[ng-repeat="item in attributeDto.settings.values"]`, "service")
	// hoeInfo.Duration = browser.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(15) > td:nth-child(2) > attribute-text-box span`, "duration")
	// hoeInfo.WorkTime = browser.GetElementText(detailInfoElement, `product-attribute table > tbody > tr:nth-child(16) > td:nth-child(2) > attribute-text-box span`, "work_time")

	// Get report urls
	// var reportUrls []string
	// page.MustElement(`li[index="2"] a.nav-link`).MustClick().MustWaitLoad()
	// time.Sleep(1 * time.Second)
	// reportTabElement := page.MustElement(`product-review[ng-if="reviewTabLoaded"]`)

	// for {
	// 	reportElements, err := reportTabElement.Elements(`div[ng-repeat="review in reviews"]`)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if len(reportElements) == 0 && len(reportUrls) == 0 {
	// 		continue
	// 		// panic(fmt.Errorf("empty reportElements ?: %v", err))
	// 	}

	// 	for _, reportElement := range reportElements {
	// 		btnElement, err := reportElement.Element(`a.view_more_report`)
	// 		if err != nil {
	// 			panic(fmt.Errorf("failed to get view_more_report: %v", err))
	// 		}

	// 		reportUrl, err := btnElement.Attribute("href")
	// 		if err != nil {
	// 			panic(fmt.Errorf("failed to get reportUrl: %v", err))
	// 		}
	// 		reportUrls = append(reportUrls, *reportUrl)
	// 	}

	// 	if len(reportUrls) == 0 {
	// 		break
	// 	}

	// 	goNextPageBtn, err := page.Timeout(1 * time.Second).Element(`product-review li.pagination-next:not(.disabled) a[ng-click]`)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		goNextPageBtn.MustClick().MustWaitLoad().CancelTimeout()
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }

	return &hoeInfo, err
}
