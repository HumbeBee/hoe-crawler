package interfaces

import "github.com/HumbeBee/hoe-crawler/internal/definitions"

type CloudflareBypasser interface {
	RequestToBypasser(url string) ([]byte, error)
	ParseResponse(rawResponse []byte) (*definitions.BypassResult, error)
}
