package interfaces

import (
	"github.com/HumbeBee/hoe-crawler/internal/dto"
)

type Scraper interface {
	GetDetailURLs(baseURL string, relativeURL string) ([]string, error)
	GetRawHoeData(baseURL string, relativeURL string) (*dto.RawHoeData, error)
}
