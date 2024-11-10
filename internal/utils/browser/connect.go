package browser

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func ConnectToPage(url string) (*rod.Browser, *rod.Page, *rod.Element, error) {
	rodBrowser := rod.New()
	if err := rodBrowser.Connect(); err != nil {
		return nil, nil, nil, err
	}

	page, err := rodBrowser.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		return nil, nil, nil, err
	}

	if err := page.WaitStable(time.Duration(30)); err != nil {
		return nil, nil, nil, err
	}

	root := page.MustElement("html")
	return rodBrowser, page, root, nil
}
