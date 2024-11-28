package interfaces

import (
	"github.com/HumbeBee/hoe-crawler/internal/dto"
)

type Scraper interface {
	GetDetailURLs() ([]string, error)
	GetRawHoeData(url string) (*dto.RawHoeData, error)
}
