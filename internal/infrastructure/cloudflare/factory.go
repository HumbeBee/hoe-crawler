package cloudflare

import "github.com/HumbeBee/hoe-crawler/internal/infrastructure/interfaces"

// internal/infrastructure/cloudflare/factory.go
func NewBypasser(bypasserType string) interfaces.CloudflareBypasser {
	switch bypasserType {
	case "yoori":
		return newYooriBypasser()
	default:
		return nil
	}
}
