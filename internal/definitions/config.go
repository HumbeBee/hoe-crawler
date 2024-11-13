package definitions

import "github.com/haovoanh28/gai-webscraper/internal/utils/logutil"

type ScraperConfig struct {
	BaseURL           string
	RequestsPerSecond float64
	Logger            *logutil.Logger
}
