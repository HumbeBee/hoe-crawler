package gaito

import (
	"log"

	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
)

type Scraper struct {
	baseURL        string
	requestsPerSec float64
	ErrorHandler   *errutil.ErrorHandler
}

func NewScraper(baseURL string, requestsPerSec float64, logger *log.Logger) *Scraper {
	return &Scraper{baseURL: baseURL, requestsPerSec: requestsPerSec, ErrorHandler: errutil.NewErrorHandler(logger, errutil.DEBUG)}
}

func (s *Scraper) GetList() ([]string, error) {
	return s.processListPage()
}

func (s *Scraper) GetDetail(url string) (*models.HoeInfo, error) {
	return s.processDetailPage(url)
}
