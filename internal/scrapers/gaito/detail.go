package gaito

import (
	"time"

	"github.com/haovoanh28/gai-webscraper/internal/transformer"

	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/browser"
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/utils"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
)

type detailPageScraper struct {
	conn *browser.Connection
	url  string
}

func newDetailPageScraper(conn *browser.Connection, url string) *detailPageScraper {
	return &detailPageScraper{conn: conn, url: url}
}

func (s *detailPageScraper) getBasicInfo() (*models.HoeInfo, error) {
	id := utils.GetIDFromUrl(s.url)

	containerEle, err := s.conn.Root.Find(detailPageSelectors.PageContainer)
	if err != nil {
		return nil, errutil.WrapError("get container element", err, s.url)
	}

	detailInfoTabEle, err := containerEle.Find(detailPageSelectors.DetailInfoTab)
	if err != nil {
		return nil, errutil.WrapError("get detail info tab element", err, s.url)
	}

	hoeInfo := &models.HoeInfo{
		OriginID: id,
		Url:      s.url,
	}

	hoeInfo.Name = containerEle.MustFind(detailPageSelectors.Name).MustGetText()
	hoeInfo.ImageUrl = containerEle.MustFind(detailPageSelectors.ImageUrl).MustGetAttribute("src")
	hoeInfo.Price = detailInfoTabEle.MustFind(detailPageSelectors.Price).MustGetText()
	hoeInfo.Phone = detailInfoTabEle.MustFind(detailPageSelectors.Phone).MustGetText()
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

	hoeInfo = transformer.TransformHoe(hoeInfo)

	return hoeInfo, nil
}

func (s *detailPageScraper) getReportURLs() ([]*models.HoeReport, error) {
	var reports []*models.HoeReport
	reportTabEle, err := s.conn.Root.Find(detailPageSelectors.ReportTab)
	if err != nil {
		return nil, errutil.WrapError("get report tab element", err, s.url)
	}

	if err := reportTabEle.Click(); err != nil {
		return nil, errutil.WrapError("click report tab element", err, s.url)
	}

	if err := reportTabEle.WaitVisible(); err != nil {
		return nil, errutil.WrapError("wait report tab element visible", err, s.url)
	}

	reportTabContentEle, err := s.conn.Root.Find(detailPageSelectors.ReportTabContent)
	if err != nil {
		return nil, errutil.WrapError("get report tab content element", err, s.url)
	}

	for {
		// reportsEle, err := browser.GetMultipleElementsWithRetry(reportTabContentEle, detailPageSelectors.ReportList)
		reportsEle, err := reportTabContentEle.FindAll(detailPageSelectors.ReportList)
		if err != nil {
			return nil, errutil.WrapError("get report elements", err, s.url)
		}

		for _, reportEle := range reportsEle {
			// reportUrl, err := browser.GetElementAttribute(reportEle, detailPageSelectors.ReportViewMoreBtn, "href")
			reportUrl, err := reportEle.MustFind(detailPageSelectors.ReportViewMoreBtn).GetAttribute("href")
			if err != nil {
				return nil, errutil.WrapError("get report url", err, s.url)
			}
			reports = append(reports, &models.HoeReport{
				ReportURL: reportUrl,
			})
		}

		goNextPageBtn, err := s.conn.Root.Find(detailPageSelectors.ReportGoNextPageBtn)
		if err != nil {
			break
		} else {
			// Click go next page button
			if err := goNextPageBtn.Click(); err != nil {
				return nil, errutil.WrapError("click go next page button", err, s.url)
			}

			time.Sleep(1 * time.Second)
		}
	}

	return reports, nil
}
