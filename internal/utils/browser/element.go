package browser

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
)

type ElementFinder interface {
	Elements(selector string) ([]*rod.Element, error)
}

func getVisibleElement(rodElement *rod.Element, selector string) (*rod.Element, error) {
	element, err := GetElementWithRetry(rodElement, selector)
	if err != nil {
		return nil, fmt.Errorf("get element: %w", err)
	}

	if err = element.WaitVisible(); err != nil {
		return nil, fmt.Errorf("wait visible: %w", err)
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
	// element, err := rodElement.Element(selector)
	element, err := getVisibleElement(rodElement, selector)
	if err != nil {
		return "", err
	}

	text, err := element.Text()
	if err != nil {
		return "", fmt.Errorf("get text: %w", err)
	}

	return strings.TrimSpace(text), nil
}

func GetElementAttribute(rodElement *rod.Element, selector string, attributeName string) (string, error) {
	element, err := getVisibleElement(rodElement, selector)
	if err != nil {
		return "", err
	}

	attribute, err := element.Attribute(attributeName)
	if err != nil {
		return "", fmt.Errorf("get attribute: %w", err)
	}

	if attribute == nil {
		return "", nil
	}

	return *attribute, nil
}

func GetElementsText(rodElement *rod.Element, selector string) (string, error) {
	elements, err := rodElement.Elements(selector)
	if err != nil {
		return "", fmt.Errorf("get elements: %w", err)
	}

	var texts []string
	for _, element := range elements {
		text, err := element.Text()
		if err != nil {
			return "", fmt.Errorf("get text: %w", err)
		}

		text = strings.TrimSpace(text)
		if text != "" {
			texts = append(texts, text)
		}
	}

	return strings.Join(texts, ", "), nil
}
