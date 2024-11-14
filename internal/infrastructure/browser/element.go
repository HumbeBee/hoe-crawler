package browser

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
)

func GetVisibleElement(rodElement *rod.Element, selector string) (*rod.Element, error) {
	element, err := GetElementWithRetry(rodElement, selector)
	if err != nil {
		return nil, err
	}

	if err = element.WaitVisible(); err != nil {
		return nil, err
	}

	return element, nil
}

func GetElementWithRetry(rodElement *rod.Element, selector string) (*rod.Element, error) {
	return retryRodElement(func() (*rod.Element, error) {
		return rodElement.Element(selector)
	})
}

func GetMultipleElementsWithRetry(rodElement *rod.Element, selector string) ([]*rod.Element, error) {
	elements, err := retryRodElement(func() ([]*rod.Element, error) {
		return rodElement.Elements(selector)
	})

	if err != nil {
		return nil, err
	}

	return elements, nil
}

func GetElementText(rodElement *rod.Element, selector string) (string, error) {
	element, err := GetVisibleElement(rodElement, selector)
	if err != nil {
		return "", err
	}

	text, err := element.Text()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}

func GetElementAttribute(rodElement *rod.Element, selector string, attributeName string) (string, error) {
	element, err := GetVisibleElement(rodElement, selector)
	if err != nil {
		return "", err
	}

	attribute, err := element.Attribute(attributeName)
	if err != nil {
		return "", err
	}

	if attribute == nil {
		return "", nil
	}

	return *attribute, nil
}

func GetElementsText(rodElement *rod.Element, selector string) (string, error) {
	elements, err := GetMultipleElementsWithRetry(rodElement, selector)
	if err != nil {
		return "", err
	}

	var texts []string
	for _, element := range elements {
		text, err := element.Text()
		if err != nil {
			return "", err
		}

		text = strings.TrimSpace(text)
		if text != "" {
			texts = append(texts, text)
		}
	}

	return strings.Join(texts, ", "), nil
}

// MUST FUNCTIONS
// Thêm mấy hàm Must để YOLO
func MustGetVisibleElement(rodElement *rod.Element, selector string) *rod.Element {
	element, err := GetVisibleElement(rodElement, selector)
	if err != nil {
		panic(fmt.Errorf("selector '%s': %w", selector, err))
	}
	return element
}

func MustGetElementText(rodElement *rod.Element, selector string) string {
	text, err := GetElementText(rodElement, selector)
	if err != nil {
		panic(fmt.Errorf("cannot get text from selector '%s': %w", selector, err))
	}
	return text
}

func MustGetElementsText(rodElement *rod.Element, selector string) string {
	text, err := GetElementsText(rodElement, selector)
	if err != nil {
		panic(fmt.Errorf("cannot get text from selector '%s': %w", selector, err))
	}
	return text
}

func MustGetElementAttribute(rodElement *rod.Element, selector string, attributeName string) string {
	attr, err := GetElementAttribute(rodElement, selector, attributeName)
	if err != nil {
		panic(fmt.Errorf("cannot get attribute '%s' from selector '%s': %w",
			attributeName, selector, err))
	}
	return attr
}
