package roddriver

import (
	"strings"

	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/interfaces"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type rodElement struct {
	element *rod.Element
}

// Consider using a factory pattern to return rodElement, especially if switching libraries later
func NewElement(element *rod.Element) interfaces.Element {
	return &rodElement{element: element}
}

func (re *rodElement) getElementWithRetry(selector string) (*rod.Element, error) {
	return retryRodElement(func() (*rod.Element, error) {
		return re.element.Element(selector)
	})
}

func (re *rodElement) getMultipleElementsWithRetry(selector string) ([]*rod.Element, error) {
	return retryRodElement(func() ([]*rod.Element, error) {
		return re.element.Elements(selector)
	})
}

// Implementing Element interface
func (re *rodElement) Find(selector string) (interfaces.Element, error) {
	elem, err := re.getElementWithRetry(selector)
	if err != nil {
		return nil, err
	}
	return NewElement(elem), nil
}

func (re *rodElement) FindAll(selector string) ([]interfaces.Element, error) {
	elems, err := re.getMultipleElementsWithRetry(selector)
	if err != nil {
		return nil, err
	}

	elements := make([]interfaces.Element, len(elems))
	for i, elem := range elems {
		elements[i] = NewElement(elem)
	}
	return elements, nil
}

func (re *rodElement) GetText() (string, error) {
	text, err := re.element.Text()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func (re *rodElement) GetAttribute(name string) (string, error) {
	attr, err := re.element.Attribute(name)
	if err != nil {
		return "", err
	}
	if attr == nil {
		return "", nil
	}
	return *attr, nil
}

func (re *rodElement) Click() error {
	return re.element.Click(proto.InputMouseButtonLeft, 1)
}

func (re *rodElement) WaitVisible() error {
	return re.element.WaitVisible()
}

// Must versions that panic on error
func (re *rodElement) MustFind(selector string) interfaces.Element {
	elem, err := re.Find(selector)
	if err != nil {
		panic(err)
	}
	return elem
}

func (re *rodElement) MustGetText() string {
	text, err := re.GetText()
	if err != nil {
		panic(err)
	}
	return text
}

func (re *rodElement) MustGetAttribute(name string) string {
	attr, err := re.GetAttribute(name)
	if err != nil {
		panic(err)
	}
	return attr
}
