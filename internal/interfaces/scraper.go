package interfaces

import "github.com/haovoanh28/gai-webscraper/internal/models"

type Scraper interface {
	ProcessListPage() ([]string, error)
	ProcessDetailPage(url string) (*models.HoeInfo, error)
}
