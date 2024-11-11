package browser

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func ConnectToPage(url string) (*rod.Browser, *rod.Page, *rod.Element, error) {
	rodBrowser := rod.New()
	if err := rodBrowser.Connect(); err != nil {
		return nil, nil, nil, err
	}

	// page, err = rodBrowser.Page(proto.TargetCreateTarget{URL: url})
	// if err != nil {
	// 	return nil, nil, nil, err
	// }

	// stealth must be used because of Cloudflare
	// But it only works sometimes
	fmt.Printf("js: %x\n\n", md5.Sum([]byte(stealth.JS)))
	page, err := stealth.Page(rodBrowser)
	if err != nil {
		return nil, nil, nil, err
	}

	if err := page.Navigate(url); err != nil {
		return nil, nil, nil, err
	}

	if err := page.WaitStable(time.Duration(30)); err != nil {
		return nil, nil, nil, err
	}

	root := page.MustElement("html")
	return rodBrowser, page, root, nil
}
