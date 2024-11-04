package utils

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func retryRodElement[T any](operation func() (T, error), fieldName string) (T, error) {
	maxAttempts := 20
	interval := 500 * time.Millisecond

	for attempt := 0; attempt < maxAttempts; attempt++ {
		result, err := operation()
		if err == nil {
			if slice, ok := any(result).([]*rod.Element); ok { // Check for multiple elements
				if len(slice) > 0 {
					return result, nil
				}
			} else if elem, ok := any(result).(*rod.Element); ok { // Check for single element
				if elem != nil {
					if visible, _ := elem.Visible(); visible {
						return result, nil
					}
				}
			}
		}
		time.Sleep(interval)
	}

	var zero T
	return zero, fmt.Errorf("failed to execute operation %s after %d attempts", fieldName, maxAttempts)
}
