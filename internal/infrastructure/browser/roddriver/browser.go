package roddriver

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/interfaces"
)

type rodBrowser struct {
	browser *rod.Browser
}

func (rb *rodBrowser) Connect() error {
	if err := rb.browser.Connect(); err != nil {
		return err
	}

	return nil
}

func (rb *rodBrowser) Close() {
	rb.browser.Close()
}

func (rb *rodBrowser) CreatePage() (interfaces.Page, error) {
	page, err := stealth.Page(rb.browser)
	if err != nil {
		return nil, err
	}

	return &rodPage{page: page}, nil
}

// Return rodBrowser, if later uses another library, consider using a factory pattern
func NewBrowser(timeout time.Duration) interfaces.Browser {
	return &rodBrowser{browser: rod.New().Timeout(timeout)}
}
