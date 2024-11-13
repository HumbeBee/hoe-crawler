package gaito

import (
	"github.com/haovoanh28/gai-webscraper/internal/definitions"
)

type Scraper struct {
	definitions.ScraperConfig
}

func NewScraper(config definitions.ScraperConfig) *Scraper {
	return &Scraper{
		config,
	}
}
