package gaito

import (
	"strings"
	"time"

	"github.com/go-rod/rod/lib/proto"
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/utils"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
)

func (s *scraper) ProcessDetailPage(detailUrl string) (*models.HoeInfo, error) {
	url := s.BaseURL + detailUrl

	id := utils.GetIDFromUrl(detailUrl)

	// Wait until content element is visible
	conn, err := browser.ConnectToPage(url, 2*time.Minute)
	if err != nil {
		return nil, errutil.WrapError("connect to page", err, url)
	}
	defer conn.Close()

	containerEle, err := browser.GetVisibleElement(conn.Root, detailPageSelectors.PageContainer)
	if err != nil {
		return nil, errutil.WrapError("get container element", err, url)
	}

	detailInfoTabEle, err := browser.GetVisibleElement(containerEle, detailPageSelectors.DetailInfoTab)
	if err != nil {
		return nil, errutil.WrapError("get detail info tab element", err, url)
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
	reportTabEle, err := browser.GetVisibleElement(conn.Root, detailPageSelectors.ReportTab)
	if err != nil {
		return nil, errutil.WrapError("get report tab element", err, url)
	}

	if err := reportTabEle.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return nil, errutil.WrapError("click report tab element", err, url)
	}

	if err := reportTabEle.WaitVisible(); err != nil {
		return nil, errutil.WrapError("wait report tab element visible", err, url)
	}

	reportTabContentEle, err := browser.GetVisibleElement(conn.Root, detailPageSelectors.ReportTabContent)
	if err != nil {
		return nil, errutil.WrapError("get report tab content element", err, url)
	}

	for {
		reportsEle, err := browser.GetMultipleElementsWithRetry(reportTabContentEle, detailPageSelectors.ReportList)
		if err != nil {
			return nil, errutil.WrapError("get report elements", err, url)
		}

		for _, reportEle := range reportsEle {
			reportUrl, err := browser.GetElementAttribute(reportEle, detailPageSelectors.ReportViewMoreBtn, "href")
			if err != nil {
				return nil, errutil.WrapError("get report url", err, url)
			}
			reports = append(reports, &models.HoeReport{
				ReportURL: reportUrl,
			})
		}

		goNextPageBtn, err := browser.GetVisibleElement(conn.Root, detailPageSelectors.ReportGoNextPageBtn)
		if err != nil {
			break
		} else {
			// Click go next page button
			if err := goNextPageBtn.Click(proto.InputMouseButtonLeft, 1); err != nil {
				return nil, errutil.WrapError("click go next page button", err, url)
			}

			time.Sleep(1 * time.Second)
		}
	}

	hoeInfo.Reports = reports
	return &hoeInfo, err
}
