package gaito

import (
	"fmt"
	"time"

	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/browser"
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
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

	conn, err := browser.ConnectToPage(url, 30*time.Second)
	if err != nil {
		return nil, errutil.WrapError("connect to page", err, url)
	}
	defer conn.Close()

	listScraper := newListPageScraper(conn, url)
	return listScraper.getHoeURLs()
}

func (s *scraper) GetRawHoeInfo(detailUrl string) (*models.HoeInfo, error) {
	url := s.BaseURL + detailUrl

	// Wait until content element is visible
	conn, err := browser.ConnectToPage(url, 2*time.Minute)
	if err != nil {
		return nil, errutil.WrapError("connect to page", err, url)
	}
	defer conn.Close()

	detailScraper := newDetailPageScraper(conn, url)
	hoeInfo, err := detailScraper.getBasicInfo()
	if err != nil {
		return nil, errutil.WrapError("get basic info", err, url)
	}

	// Get report urls
	// reports, err := detailScraper.getReportURLs()
	// if err != nil {
	// 	return nil, errutil.WrapError("get report urls", err, url)
	// }

	// hoeInfo.Reports = reports
	return hoeInfo, nil
}
