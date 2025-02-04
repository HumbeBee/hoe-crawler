package definitions

import (
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type ScraperConfig struct {
	SiteID   uint
	SiteName string
	Logger   *logutil.Logger
}
