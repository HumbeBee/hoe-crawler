package definitions

import (
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type ScraperConfig struct {
	SiteID            uint
	SiteName          string
	RequestsPerSecond float64
	Logger            *logutil.Logger
}
