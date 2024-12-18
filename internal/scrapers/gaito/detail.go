package gaito

import (
	"fmt"
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/dto"

	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
	"github.com/HumbeBee/hoe-crawler/internal/utils"
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
		return nil, fmt.Errorf("get container element: %w", err)
	}

	detailInfoTabEle, err := containerEle.Find(detailPageSelectors.DetailInfoTab)
	if err != nil {
		return nil, fmt.Errorf("get detail info tab element: %w", err)
	}

	rawInfo := &dto.RawHoeData{
		SiteID:   s.siteID,
		Url:      s.url,
		OriginID: id,
	}

	rawInfo.CityName, err = containerEle.FindAndGetText(detailPageSelectors.CityName)
	if err != nil {
		return nil, fmt.Errorf("get city name: %w", err)
	}

	rawInfo.DistrictName, err = containerEle.FindAndGetText(detailPageSelectors.DistrictName)
	if err != nil {
		return nil, fmt.Errorf("get district name: %w", err)
	}

	rawInfo.Name, err = containerEle.FindAndGetText(detailPageSelectors.Name)
	if err != nil {
		return nil, fmt.Errorf("get name: %w", err)
	}

	rawInfo.ImageUrl, err = containerEle.FindAndGetAttribute(detailPageSelectors.ImageUrl, "src")
	if err != nil {
		return nil, fmt.Errorf("get image URL: %w", err)
	}

	rawInfo.Price, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Price)
	if err != nil {
		return nil, fmt.Errorf("get price: %w", err)
	}

	rawInfo.Phone, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Phone)
	if err != nil {
		return nil, fmt.Errorf("get phone: %w", err)
	}

	rawInfo.Area, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Address)
	if err != nil {
		return nil, fmt.Errorf("get area: %w", err)
	}

	rawInfo.Provider, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Author)
	if err != nil {
		return nil, fmt.Errorf("get provider: %w", err)
	}

	rawInfo.Status, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Status)
	if err != nil {
		return nil, fmt.Errorf("get status: %w", err)
	}

	rawInfo.BirthYear, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.BirthYear)
	if err != nil {
		return nil, fmt.Errorf("get birth year: %w", err)
	}

	rawInfo.Height, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Height)
	if err != nil {
		return nil, fmt.Errorf("get height: %w", err)
	}

	rawInfo.Weight, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Weight)
	if err != nil {
		return nil, fmt.Errorf("get weight: %w", err)
	}

	rawInfo.Country, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Country)
	if err != nil {
		return nil, fmt.Errorf("get country: %w", err)
	}

	rawInfo.Service, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Service)
	if err != nil {
		return nil, fmt.Errorf("get service: %w", err)
	}

	rawInfo.Duration, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.Duration)
	if err != nil {
		return nil, fmt.Errorf("get duration: %w", err)
	}

	rawInfo.WorkTime, err = detailInfoTabEle.FindAndGetText(detailPageSelectors.WorkTime)
	if err != nil {
		return nil, fmt.Errorf("get work time: %w", err)
	}

	return rawInfo, nil
}

func (s *detailPageScraper) getReportURLs() ([]string, error) {
	var reports []string
	reportTabEle, err := s.conn.Root.Find(detailPageSelectors.ReportTab)
	if err != nil {
		return nil, fmt.Errorf("get report tab element: %w", err)
	}

	if err := reportTabEle.Click(); err != nil {
		return nil, fmt.Errorf("click report tab element: %w", err)
	}

	if err := reportTabEle.WaitElement(); err != nil {
		return nil, fmt.Errorf("wait report tab element visible: %w", err)
	}

	reportTabContentEle, err := s.conn.Root.Find(detailPageSelectors.ReportTabContent)
	if err != nil {
		return nil, fmt.Errorf("get report tab content element: %w", err)
	}

	for {
		// reportsEle, err := browser.GetMultipleElementsWithRetry(reportTabContentEle, detailPageSelectors.ReportList)
		reportsEle, err := reportTabContentEle.FindAll(detailPageSelectors.ReportList)
		if err != nil {
			return nil, fmt.Errorf("get report elements: %w", err)
		}

		for _, reportEle := range reportsEle {
			// reportUrl, err := browser.GetElementAttribute(reportEle, detailPageSelectors.ReportViewMoreBtn, "href")
			reportUrl, err := reportEle.MustFind(detailPageSelectors.ReportViewMoreBtn).GetAttribute("href")
			if err != nil {
				return nil, fmt.Errorf("get report url: %w", err)
			}
			reports = append(reports, reportUrl)
		}

		goNextPageBtn, err := s.conn.Root.Find(detailPageSelectors.ReportGoNextPageBtn)
		if err != nil {
			break
		} else {
			// Click go next page button
			if err := goNextPageBtn.Click(); err != nil {
				return nil, fmt.Errorf("click go next page button: %w", err)
			}

			time.Sleep(1 * time.Second)
		}
	}

	return reports, nil
}
