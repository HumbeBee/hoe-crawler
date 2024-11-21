package roddriver

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/interfaces"
)

type rodPage struct {
	page *rod.Page
}

func (rp *rodPage) Navigate(url string) error {
	return rp.page.Navigate(url)
}

func (rp *rodPage) WaitStable(timeout time.Duration) error {
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
