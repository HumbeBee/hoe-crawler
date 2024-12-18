package gaito

import (
	"fmt"
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

func (s *scraper) GetDetailURLs(baseURL string, relativeURL string) ([]string, error) {
	fullURL := baseURL + relativeURL
	s.Logger.Info(fmt.Sprintf("Processing %s", fullURL))

	conn, err := browser.ConnectToPage(fullURL, 5*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("connect to page %s: %w", fullURL, err)
	}
	defer conn.Close()

	listScraper := newListPageScraper(conn, fullURL)
	return listScraper.getHoeURLs()
}

func (s *scraper) GetRawHoeData(baseURL string, relativeURL string) (*dto.RawHoeData, error) {
	fullURL := baseURL + relativeURL

	// Wait until content element is visible
	conn, err := browser.ConnectToPage(fullURL, 5*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("connect to detail page %s: %w", fullURL, err)
	}
	defer conn.Close()

	detailScraper := newDetailPageScraper(conn, relativeURL, s.SiteID)
	hoeInfo, err := detailScraper.getBasicInfo()
	if err != nil {
		return nil, fmt.Errorf("get basic info: %w", err)
	}

	// Get report urls
	reports, err := detailScraper.getReportURLs()
	if err != nil {
		return nil, fmt.Errorf("get report urls: %w", err)
	}

	hoeInfo.Reports = reports

	s.Logger.Info("Raw hoe data: %+v", hoeInfo)

	return hoeInfo, nil
}
