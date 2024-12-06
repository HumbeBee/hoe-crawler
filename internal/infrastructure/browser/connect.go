package browser

import (
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

	page, err := browser.CreatePage(bypassResult.UserAgent)
	if err != nil {
		return nil, err
	}

	if err := page.Navigate(url); err != nil {
		return nil, err
	}

	if err := page.WaitStable(time.Duration(30)); err != nil {
		return nil, err
	}

	root := page.GetRootElement()
	if err := root.WaitVisible(); err != nil {
		page.Close()
		browser.Close()
		return nil, err
	}

	return &Connection{Browser: browser, Page: page, Root: root}, nil
}
