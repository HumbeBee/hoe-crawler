package browser

import (
	"github.com/HumbeBee/hoe-crawler/internal/utils"
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/interfaces"
)

type Connection struct {
	Browser interfaces.Browser
	Page    interfaces.Page
	Root    interfaces.Element
}

func (c *Connection) Close() {
	if c == nil {
		return
	}

	if c.Browser != nil {
		c.Browser.Close()
	}
	if c.Page != nil {
		c.Page.Close()
	}
}

func ConnectToPage(url string, timeout time.Duration) (*Connection, error) {
	browser := NewBrowser(RodDriver, timeout)
	if err := browser.Connect(); err != nil {
		return nil, err
	}

	bypassResult, err := browser.BypassCloudflare(url)
	if err != nil {
		return nil, err
	}

	// Since bypassing cloudflare is considered a request,
	// we should wait a bit before connecting to the page.
	utils.RandomDelay()

	page, err := browser.CreatePage(bypassResult.UserAgent)
	if err != nil {
		return nil, err
	}

	if err := page.Navigate(url); err != nil {
		return nil, err
	}

	if err := page.WaitPageLoad(time.Duration(30)); err != nil {
		return nil, err
	}

	root := page.GetRootElement()
	if err := root.WaitElement(); err != nil {
		page.Close()
		browser.Close()
		return nil, err
	}

	return &Connection{Browser: browser, Page: page, Root: root}, nil
}
