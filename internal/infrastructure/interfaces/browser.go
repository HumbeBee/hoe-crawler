package interfaces

import (
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"time"
)

type Browser interface {
	Connect() error
	BypassCloudflare(url string) (*definitions.BypassResult, error)
	CreatePage(userAgent string) (Page, error)
	Close()
}

type Page interface {
	Navigate(url string) error
	WaitStable(timeout time.Duration) error
	WaitElementsMoreThan(selector string, count int) error
	GetRootElement() Element
	Close()
}

type Element interface {
	Find(selector string) (Element, error)
	FindAll(selector string) ([]Element, error)
	GetText() (string, error)
	GetAttribute(name string) (string, error)
	FindAndGetText(selector string) (string, error)
	FindAndGetAttribute(selector, attr string) (string, error)
	Click() error
	WaitVisible() error
	MustFind(selector string) Element
	MustGetText() string
	MustGetAttribute(name string) string
}
