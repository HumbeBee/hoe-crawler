package interfaces

import "time"

type Browser interface {
	Connect() error
	BypassCloudflare(url string) error
	CreatePage() (Page, error)
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
	Click() error
	WaitVisible() error
	MustFind(selector string) Element
	MustGetText() string
	MustGetAttribute(name string) string
}
