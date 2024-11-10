package gaito

import (
	"strings"
	"time"

	"github.com/go-rod/rod/lib/proto"
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

	price := browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Price)
	// Ex: "300 k" => "300k"
	hoeInfo.Price = strings.ReplaceAll(price, "\u00A0", "")
	// Ex: "1.000k" => "1000k"
	hoeInfo.Price = strings.ReplaceAll(hoeInfo.Price, ".", "")

	// Ex: "0123.456.789" -> "0123456789"
	phone := browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Phone)
	hoeInfo.Phone = strings.ReplaceAll(phone, ".", "")

	hoeInfo.Address = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Address)
	hoeInfo.Provider = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Author)
	hoeInfo.Status = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Status)

	hoeInfo.BirthYear = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.BirthYear)
	hoeInfo.Height = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Height)
	hoeInfo.Weight = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Weight)
	hoeInfo.From = browser.MustGetElementsText(detailInfoTabEle, detailPageSelectors.From)
	hoeInfo.Service = browser.MustGetElementsText(detailInfoTabEle, detailPageSelectors.Service)
	hoeInfo.Duration = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Duration)
	hoeInfo.WorkTime = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.WorkTime)

	// Get report urls
	var reportUrls []string
	reportTabEle, err := browser.GetVisibleElement(root, detailPageSelectors.ReportTab)
	if err != nil {
		return nil, s.ErrorHandler.WrapError("get report tab element", err, url)
	}

	if err := reportTabEle.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return nil, s.ErrorHandler.WrapError("click report tab element", err, url)
	}

	if err := reportTabEle.WaitVisible(); err != nil {
		return nil, s.ErrorHandler.WrapError("wait report tab element visible", err, url)
	}

	reportTabContentEle, err := browser.GetVisibleElement(root, detailPageSelectors.ReportTabContent)
	if err != nil {
		return nil, s.ErrorHandler.WrapError("get report tab content element", err, url)
	}

	for {
		reportsEle, err := browser.GetMultipleElementsWithRetry(reportTabContentEle, detailPageSelectors.ReportList)
		if err != nil {
			return nil, s.ErrorHandler.WrapError("get report elements", err, url)
		}

		for _, reportEle := range reportsEle {
			reportUrl, err := browser.GetElementAttribute(reportEle, detailPageSelectors.ReportViewMoreBtn, "href")
			if err != nil {
				return nil, s.ErrorHandler.WrapError("get report url", err, url)
			}
			reportUrls = append(reportUrls, reportUrl)
		}

		goNextPageBtn, err := browser.GetVisibleElement(root, detailPageSelectors.ReportGoNextPageBtn)
		if err != nil {
			break
		} else {
			// Click go next page button
			if err := goNextPageBtn.Click(proto.InputMouseButtonLeft, 1); err != nil {
				return nil, s.ErrorHandler.WrapError("click go next page button", err, url)
			}

			time.Sleep(1 * time.Second)
		}
	}

	hoeInfo.ReportUrls = reportUrls
	return &hoeInfo, err
}
