package roddriver

import (
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/interfaces"
	"github.com/go-rod/rod"
)

type rodPage struct {
	page *rod.Page
}

func (rp *rodPage) Navigate(url string) error {
	return rp.page.Navigate(url)
}

func (rp *rodPage) WaitPageLoad(timeout time.Duration) error {
	if err := rp.page.WaitDOMStable(timeout, 0.1); err != nil {
		return err
	}

	return rp.page.WaitStable(timeout)
}

func (rp *rodPage) GetRootElement() interfaces.Element {
	return NewElement(rp.page.MustElement("html"))
}

func (rp *rodPage) Close() {
	rp.page.Close()
}

func (rp *rodPage) WaitElementsMoreThan(selector string, count int) error {
	return rp.page.WaitElementsMoreThan(selector, count)
}
