package gaito

import (
	"github.com/haovoanh28/gai-webscraper/internal/definitions"
)

type scraper struct {
	definitions.ScraperConfig
}

func NewScraper(config definitions.ScraperConfig) *scraper {
	return &scraper{
		config,
	}
}
