package browser

import (
	"strings"

	"github.com/go-rod/rod"
	"github.com/haovoanh28/gai-webscraper/internal/utils"
)

func GetElementWithRetry(rodElement *rod.Element, selector string, fieldName string) (*rod.Element, error) {
	return retryRodElement(func() (*rod.Element, error) {
		return rodElement.Element(selector)
	}, fieldName)
}

func GetElementText(rodElement *rod.Element, selector string, fieldName string) string {
	// element, err := rodElement.Element(selector)
	element, err := GetElementWithRetry(rodElement, selector, fieldName)
	utils.HandleError(err, "get element", fieldName)

	err = element.WaitVisible()
	utils.HandleError(err, "wait visible", fieldName)

	text, err := element.Text()
	utils.HandleError(err, "get text", fieldName)

	return strings.TrimSpace(text)
}

func GetElementAttribute(rodElement *rod.Element, selector string, attributeName string, fieldName string) string {
	element, err := GetElementWithRetry(rodElement, selector, fieldName)
	utils.HandleError(err, "get element", fieldName)

	attribute, err := element.Attribute(attributeName)
	utils.HandleError(err, "get attribute", fieldName)

	return *attribute
}

func GetElementsText(rodElement *rod.Element, selector string, fieldName string) string {
	elements, err := rodElement.Elements(selector)
	utils.HandleError(err, "get elements", fieldName)

	var texts []string
	for _, element := range elements {
		text, err := element.Text()
		utils.HandleError(err, "get text", fieldName)

		text = strings.TrimSpace(text)
		if text != "" {
			texts = append(texts, text)
		}
	}

	return strings.Join(texts, ", ")
}
