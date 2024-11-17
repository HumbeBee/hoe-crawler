package transformer

import (
	"strings"

	"github.com/haovoanh28/gai-webscraper/internal/models"
)

func TransformHoe(hoe *models.HoeInfo) *models.HoeInfo {
	hoe.Price = transformPrice(hoe.Price)
	hoe.Phone = transformPhone(hoe.Phone)

	return hoe
}

func transformPrice(price string) string {
	// First clean up any special characters and spaces
	price = strings.ReplaceAll(price, "\u00A0", "")
	price = strings.ReplaceAll(price, ",", "")
	price = strings.ToLower(price)

	// Remove any spaces
	price = strings.ReplaceAll(price, " ", "")

	// Remove 'k' and 'vnd'
	price = strings.ReplaceAll(price, "k", "")
	price = strings.ReplaceAll(price, "vnd", "")

	return price
}

func transformPhone(phone string) string {
	phone = strings.ReplaceAll(phone, ".", "")
	return phone
}
