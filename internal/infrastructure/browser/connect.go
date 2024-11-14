package browser

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

type BrowserConnection struct {
	Browser *rod.Browser
	Page    *rod.Page
	Root    *rod.Element
}

func (c *BrowserConnection) Close() {
	if c.Browser != nil {
		c.Browser.Close()
	}
	if c.Page != nil {
		c.Page.Close()
	}
}

func ConnectToPage(url string, timeout time.Duration) (*BrowserConnection, error) {
	rodBrowser := rod.New().Timeout(timeout)
	if err := rodBrowser.Connect(); err != nil {
		return nil, err
	}

	// stealth must be used because of Cloudflare
	// But it only works sometimes
	fmt.Printf("js: %x\n\n", md5.Sum([]byte(stealth.JS)))
	page, err := stealth.Page(rodBrowser)
	if err != nil {
		return nil, err
	}

	if err := page.Navigate(url); err != nil {
		return nil, err
	}

	if err := page.WaitStable(time.Duration(30)); err != nil {
		return nil, err
	}

	root := page.MustElement("html")
	return &BrowserConnection{Browser: rodBrowser, Page: page, Root: root}, nil
}
