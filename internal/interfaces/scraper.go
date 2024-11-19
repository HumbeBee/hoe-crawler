package interfaces

import (
	"github.com/haovoanh28/gai-webscraper/internal/dto"
)

type Scraper interface {
	GetDetailURLs() ([]string, error)
	GetRawHoeData(url string) (*dto.RawHoeData, error)
}
