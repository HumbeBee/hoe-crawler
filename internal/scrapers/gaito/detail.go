package gaito

import (
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/dto"

	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
	"github.com/HumbeBee/hoe-crawler/internal/utils"
	"github.com/HumbeBee/hoe-crawler/internal/utils/errutil"
)

type detailPageScraper struct {
	siteID uint
	conn   *browser.Connection
	url    string
}

func newDetailPageScraper(conn *browser.Connection, url string, siteID uint) *detailPageScraper {
	return &detailPageScraper{conn: conn, url: url, siteID: siteID}
}

func (s *detailPageScraper) getBasicInfo() (*dto.RawHoeData, error) {
	id := utils.GetIDFromUrl(s.url)

	containerEle, err := s.conn.Root.Find(detailPageSelectors.PageContainer)
	if err != nil {
		return nil, errutil.WrapError("get container element", err)
	}

	detailInfoTabEle, err := containerEle.Find(detailPageSelectors.DetailInfoTab)
	if err != nil {
		return nil, errutil.WrapError("get detail info tab element", err)
	}

	rawInfo := &dto.RawHoeData{
		SiteID:   s.siteID,
		Url:      s.url,
		OriginID: id,
	}

	rawInfo.CityName = containerEle.MustFind(detailPageSelectors.CityName).MustGetText()
	rawInfo.DistrictName = containerEle.MustFind(detailPageSelectors.DistrictName).MustGetText()
	rawInfo.Name = containerEle.MustFind(detailPageSelectors.Name).MustGetText()
	rawInfo.ImageUrl = containerEle.MustFind(detailPageSelectors.ImageUrl).MustGetAttribute("src")
	rawInfo.Price = detailInfoTabEle.MustFind(detailPageSelectors.Price).MustGetText()
	rawInfo.Phone = detailInfoTabEle.MustFind(detailPageSelectors.Phone).MustGetText()
	rawInfo.Area = detailInfoTabEle.MustFind(detailPageSelectors.Address).MustGetText()
	rawInfo.Provider = detailInfoTabEle.MustFind(detailPageSelectors.Author).MustGetText()
	rawInfo.Status = detailInfoTabEle.MustFind(detailPageSelectors.Status).MustGetText()
	rawInfo.BirthYear = detailInfoTabEle.MustFind(detailPageSelectors.BirthYear).MustGetText()
	rawInfo.Height = detailInfoTabEle.MustFind(detailPageSelectors.Height).MustGetText()
	rawInfo.Weight = detailInfoTabEle.MustFind(detailPageSelectors.Weight).MustGetText()
	rawInfo.Country = detailInfoTabEle.MustFind(detailPageSelectors.Country).MustGetText()
	rawInfo.Service = detailInfoTabEle.MustFind(detailPageSelectors.Service).MustGetText()
	rawInfo.Duration = detailInfoTabEle.MustFind(detailPageSelectors.Duration).MustGetText()
	rawInfo.WorkTime = detailInfoTabEle.MustFind(detailPageSelectors.WorkTime).MustGetText()

	return rawInfo, nil
}

func (s *detailPageScraper) getReportURLs() ([]string, error) {
	var reports []string
	reportTabEle, err := s.conn.Root.Find(detailPageSelectors.ReportTab)
	if err != nil {
		return nil, errutil.WrapError("get report tab element", err)
	}

	if err := reportTabEle.Click(); err != nil {
		return nil, errutil.WrapError("click report tab element", err)
	}

	if err := reportTabEle.WaitVisible(); err != nil {
		return nil, errutil.WrapError("wait report tab element visible", err)
	}

	reportTabContentEle, err := s.conn.Root.Find(detailPageSelectors.ReportTabContent)
	if err != nil {
		return nil, errutil.WrapError("get report tab content element", err)
	}

	for {
		// reportsEle, err := browser.GetMultipleElementsWithRetry(reportTabContentEle, detailPageSelectors.ReportList)
		reportsEle, err := reportTabContentEle.FindAll(detailPageSelectors.ReportList)
		if err != nil {
			return nil, errutil.WrapError("get report elements", err)
		}

		for _, reportEle := range reportsEle {
			// reportUrl, err := browser.GetElementAttribute(reportEle, detailPageSelectors.ReportViewMoreBtn, "href")
			reportUrl, err := reportEle.MustFind(detailPageSelectors.ReportViewMoreBtn).GetAttribute("href")
			if err != nil {
				return nil, errutil.WrapError("get report url", err)
			}
			reports = append(reports, reportUrl)
		}

		goNextPageBtn, err := s.conn.Root.Find(detailPageSelectors.ReportGoNextPageBtn)
		if err != nil {
			break
		} else {
			// Click go next page button
			if err := goNextPageBtn.Click(); err != nil {
				return nil, errutil.WrapError("click go next page button", err)
			}

			time.Sleep(1 * time.Second)
		}
	}

	return reports, nil
}
