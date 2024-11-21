package browser

import (
	"time"

	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/browser/roddriver"
	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/interfaces"
)

type DriverType string

const (
	RodDriver DriverType = "rod"
	// Future drivers can be added here
	// SeleniumDriver DriverType = "selenium"
)

// NewBrowser creates a new browser instance based on the driver type
func NewBrowser(driverType DriverType, timeout time.Duration) interfaces.Browser {
	switch driverType {
	case RodDriver:
		return roddriver.NewBrowser(timeout)
	default:
		// For now, default to Rod
		return roddriver.NewBrowser(timeout)
	}
}
