package gaito

import "github.com/haovoanh28/gai-webscraper/internal/models"

type Scraper struct {
	baseURL        string
	requestsPerSec float64
}

func NewScraper(baseURL string, requestsPerSec float64) *Scraper {
	return &Scraper{baseURL: baseURL, requestsPerSec: requestsPerSec}
}

func (s *Scraper) GetList() []string {
	return s.processListPage()
}

func (s *Scraper) GetDetail(url string) *models.HoeInfo {
	result := ProcessDetailPage(url)
	return &result
}
