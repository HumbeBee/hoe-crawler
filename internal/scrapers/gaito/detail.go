package gaito

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod/lib/proto"
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/utils"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
)

func (s *gaitoScraper) ProcessDetailPage(detailUrl string) (*models.HoeInfo, error) {
	url := s.BaseURL + detailUrl

	id := utils.GetIDFromUrl(detailUrl)

	// Wait until content element is visible
	rodBrowser, page, root, err := browser.ConnectToPage(url)
	if err != nil {
		return nil, fmt.Errorf("connect to page ", err, url)
	}
	defer rodBrowser.Close()
	defer page.Close()

	containerEle, err := browser.GetVisibleElement(root, detailPageSelectors.PageContainer)
	if err != nil {
		return nil, fmt.Errorf("get container element at %s (Cloudflare block ?): %w", url, err)
	}

	detailInfoTabEle, err := browser.GetVisibleElement(containerEle, detailPageSelectors.DetailInfoTab)
	if err != nil {
		return nil, fmt.Errorf("get detail info tab element at %v: %w", url, err)
	}

	hoeInfo := models.HoeInfo{
		OriginID: id,
		Url:      url,
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
	hoeInfo.Country = browser.MustGetElementsText(detailInfoTabEle, detailPageSelectors.Country)
	hoeInfo.Service = browser.MustGetElementsText(detailInfoTabEle, detailPageSelectors.Service)
	hoeInfo.Duration = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.Duration)
	hoeInfo.WorkTime = browser.MustGetElementText(detailInfoTabEle, detailPageSelectors.WorkTime)

	// Get report urls
	var reports []*models.HoeReport
	reportTabEle, err := browser.GetVisibleElement(root, detailPageSelectors.ReportTab)
	if err != nil {
		return nil, fmt.Errorf("get report tab element at %s: %w", url, err)
	}

	if err := reportTabEle.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return nil, fmt.Errorf("click report tab element at %s: %w", url, err)
	}

	if err := reportTabEle.WaitVisible(); err != nil {
		return nil, fmt.Errorf("wait report tab element visible at %s: %w", url, err)
	}

	reportTabContentEle, err := browser.GetVisibleElement(root, detailPageSelectors.ReportTabContent)
	if err != nil {
		return nil, fmt.Errorf("get report tab content element at %s: %w", url, err)
	}

	for {
		reportsEle, err := browser.GetMultipleElementsWithRetry(reportTabContentEle, detailPageSelectors.ReportList)
		if err != nil {
			return nil, fmt.Errorf("get report elements at %s: %w", url, err)
		}

		for _, reportEle := range reportsEle {
			reportUrl, err := browser.GetElementAttribute(reportEle, detailPageSelectors.ReportViewMoreBtn, "href")
			if err != nil {
				return nil, fmt.Errorf("get report url at %s: %w", url, err)
			}
			reports = append(reports, &models.HoeReport{
				ReportURL: reportUrl,
			})
		}

		goNextPageBtn, err := browser.GetVisibleElement(root, detailPageSelectors.ReportGoNextPageBtn)
		if err != nil {
			break
		} else {
			// Click go next page button
			if err := goNextPageBtn.Click(proto.InputMouseButtonLeft, 1); err != nil {
				// return nil, s.ErrorHandler.WrapError("click go next page button", err, url)
				return nil, fmt.Errorf("click go next page button at %s: %w", url, err)
			}

			time.Sleep(1 * time.Second)
		}
	}

	hoeInfo.Reports = reports
	return &hoeInfo, err
}
