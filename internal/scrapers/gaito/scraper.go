package gaito

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/utils/errutil"
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/dto"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
)

type scraper struct {
	definitions.ScraperConfig
}

func NewScraper(config definitions.ScraperConfig) *scraper {
	return &scraper{
		config,
	}
}

func (s *scraper) GetDetailURLs() ([]string, error) {
	url := s.BaseURL + "/gai-goi/khu-vuc/Hồ%20Chí%20Minh/Quận%207"

	s.Logger.Info(fmt.Sprintf("Processing %s", url))

	conn, err := browser.ConnectToPage(url, 5*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("connect to page %s: %w", url, err)
	}
	defer conn.Close()

	listScraper := newListPageScraper(conn, url)
	return listScraper.getHoeURLs()
}

func (s *scraper) GetRawHoeData(detailUrl string) (*dto.RawHoeData, error) {
	url := s.BaseURL + detailUrl

	// Wait until content element is visible
	conn, err := browser.ConnectToPage(url, 5*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("connect to detail page %s: %w", url, err)
	}
	defer conn.Close()

	detailScraper := newDetailPageScraper(conn, url, s.SiteID)
	hoeInfo, err := detailScraper.getBasicInfo()
	if err != nil {
		return nil, fmt.Errorf("get basic info %s: %w", url, err)
	}

	// Get report urls
	reports, err := detailScraper.getReportURLs()
	if err != nil {
		return nil, errutil.WrapError("get report urls", err, url)
	}

	hoeInfo.Reports = reports
	return hoeInfo, nil
}
