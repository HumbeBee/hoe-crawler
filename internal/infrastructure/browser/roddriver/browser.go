package roddriver

import (
	"github.com/go-rod/rod/lib/proto"
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/interfaces"
	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
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

func (rb *rodBrowser) CreatePage(userAgent string) (interfaces.Page, error) {
	page, err := stealth.Page(rb.browser)
	if err != nil {
		return nil, err
	}

	err = page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: userAgent,
	})

	return &rodPage{page: page}, nil
}

// Return rodBrowser, if later uses another library, consider using a factory pattern
func NewBrowser(timeout time.Duration) interfaces.Browser {
	return &rodBrowser{browser: rod.New().Timeout(timeout)}
}
