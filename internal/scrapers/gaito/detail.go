package gaito

import (
	"strings"
	"time"

	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/browser"
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/utils"
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

	// containerEle, err := browser.GetVisibleElement(conn.Root, detailPageSelectors.PageContainer)
	containerEle, err := conn.Root.Find(detailPageSelectors.PageContainer)
	if err != nil {
		return nil, errutil.WrapError("get container element", err, url)
	}

	// detailInfoTabEle, err := browser.GetVisibleElement(containerEle, detailPageSelectors.DetailInfoTab)
	detailInfoTabEle, err := containerEle.Find(detailPageSelectors.DetailInfoTab)
	if err != nil {
		return nil, errutil.WrapError("get detail info tab element", err, url)
	}

	hoeInfo := models.HoeInfo{
		OriginID: id,
		Url:      url,
	}

	hoeInfo.Name = containerEle.MustFind(detailPageSelectors.Name).MustGetText()

	hoeInfo.ImageUrl = containerEle.MustFind(detailPageSelectors.ImageUrl).MustGetAttribute("src")

	price := detailInfoTabEle.MustFind(detailPageSelectors.Price).MustGetText()
	// Ex: "300 k" => "300k"
	hoeInfo.Price = strings.ReplaceAll(price, "\u00A0", "")
	// Ex: "1.000k" => "1000k"
	hoeInfo.Price = strings.ReplaceAll(hoeInfo.Price, ".", "")

	// Ex: "0123.456.789" -> "0123456789"
	phone := detailInfoTabEle.MustFind(detailPageSelectors.Phone).MustGetText()
	hoeInfo.Phone = strings.ReplaceAll(phone, ".", "")

	hoeInfo.Address = detailInfoTabEle.MustFind(detailPageSelectors.Address).MustGetText()
	hoeInfo.Provider = detailInfoTabEle.MustFind(detailPageSelectors.Author).MustGetText()
	hoeInfo.Status = detailInfoTabEle.MustFind(detailPageSelectors.Status).MustGetText()

	hoeInfo.BirthYear = detailInfoTabEle.MustFind(detailPageSelectors.BirthYear).MustGetText()
	hoeInfo.Height = detailInfoTabEle.MustFind(detailPageSelectors.Height).MustGetText()
	hoeInfo.Weight = detailInfoTabEle.MustFind(detailPageSelectors.Weight).MustGetText()
	hoeInfo.Country = detailInfoTabEle.MustFind(detailPageSelectors.Country).MustGetText()
	hoeInfo.Service = detailInfoTabEle.MustFind(detailPageSelectors.Service).MustGetText()
	hoeInfo.Duration = detailInfoTabEle.MustFind(detailPageSelectors.Duration).MustGetText()
	hoeInfo.WorkTime = detailInfoTabEle.MustFind(detailPageSelectors.WorkTime).MustGetText()

	// Get report urls
	var reports []*models.HoeReport
	reportTabEle, err := conn.Root.Find(detailPageSelectors.ReportTab)
	if err != nil {
		return nil, errutil.WrapError("get report tab element", err, url)
	}

	if err := reportTabEle.Click(); err != nil {
		return nil, errutil.WrapError("click report tab element", err, url)
	}

	if err := reportTabEle.WaitVisible(); err != nil {
		return nil, errutil.WrapError("wait report tab element visible", err, url)
	}

	reportTabContentEle, err := conn.Root.Find(detailPageSelectors.ReportTabContent)
	if err != nil {
		return nil, errutil.WrapError("get report tab content element", err, url)
	}

	for {
		// reportsEle, err := browser.GetMultipleElementsWithRetry(reportTabContentEle, detailPageSelectors.ReportList)
		reportsEle, err := reportTabContentEle.FindAll(detailPageSelectors.ReportList)
		if err != nil {
			return nil, errutil.WrapError("get report elements", err, url)
		}

		for _, reportEle := range reportsEle {
			// reportUrl, err := browser.GetElementAttribute(reportEle, detailPageSelectors.ReportViewMoreBtn, "href")
			reportUrl, err := reportEle.MustFind(detailPageSelectors.ReportViewMoreBtn).GetAttribute("href")
			if err != nil {
				return nil, errutil.WrapError("get report url", err, url)
			}
			reports = append(reports, &models.HoeReport{
				ReportURL: reportUrl,
			})
		}

		goNextPageBtn, err := conn.Root.Find(detailPageSelectors.ReportGoNextPageBtn)
		if err != nil {
			break
		} else {
			// Click go next page button
			if err := goNextPageBtn.Click(); err != nil {
				return nil, errutil.WrapError("click go next page button", err, url)
			}

			time.Sleep(1 * time.Second)
		}
	}

	hoeInfo.Reports = reports

	return &hoeInfo, err
}
