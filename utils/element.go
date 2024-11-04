package utils

import (
	"strings"

	"github.com/go-rod/rod"
)

func GetElementWithRetry(rodElement *rod.Element, selector string, fieldName string) (*rod.Element, error) {
	return retryRodElement(func() (*rod.Element, error) {
		return rodElement.Element(selector)
	}, fieldName)
}

func GetElementText(rodElement *rod.Element, selector string, fieldName string) string {
	// element, err := rodElement.Element(selector)
	element, err := GetElementWithRetry(rodElement, selector, fieldName)
	HandleError(err, "get element", fieldName)

	err = element.WaitVisible()
	HandleError(err, "wait visible", fieldName)

	text, err := element.Text()
	HandleError(err, "get text", fieldName)

	return strings.TrimSpace(text)
}

func GetElementAttribute(rodElement *rod.Element, selector string, attributeName string, fieldName string) string {
	element, err := GetElementWithRetry(rodElement, selector, fieldName)
	HandleError(err, "get element", fieldName)

	attribute, err := element.Attribute(attributeName)
	HandleError(err, "get attribute", fieldName)

	return *attribute
}

func GetElementsText(rodElement *rod.Element, selector string, fieldName string) string {
	elements, err := rodElement.Elements(selector)
	HandleError(err, "get elements", fieldName)

	var texts []string
	for _, element := range elements {
		text, err := element.Text()
		HandleError(err, "get text", fieldName)

		text = strings.TrimSpace(text)
		if text != "" {
			texts = append(texts, text)
		}
	}

	return strings.Join(texts, ", ")
}
