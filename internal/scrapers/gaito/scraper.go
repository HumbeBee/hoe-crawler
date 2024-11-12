package gaito

import (
	"github.com/haovoanh28/gai-webscraper/internal/definitions"
)

type gaitoScraper struct {
	definitions.ScraperConfig
}

func NewScraper(config definitions.ScraperConfig) *gaitoScraper {
	return &gaitoScraper{
		config,
	}
}
