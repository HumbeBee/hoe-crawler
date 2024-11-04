package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func GetIDFromUrl(url string) string {
	url = strings.TrimSuffix(url, "/")

	// Pattern 1: "/x/y/id/z" - extract ID between segments
	re1 := regexp.MustCompile(`/([^/]+)/[^/]+$`)
	if match := re1.FindStringSubmatch(url); len(match) >= 2 {
		return match[1]
	}

	// Pattern 2: "/a/id" - extract ID at the end
	re2 := regexp.MustCompile(`/([^/]+)$`)
	if match := re2.FindStringSubmatch(url); len(match) >= 2 {
		return match[1]
	}

	return ""
}

func HandleError(err error, operation string, fieldName string) {
	if err != nil {
		panic(fmt.Errorf(`failed to %s "%s": %v`, operation, fieldName, err))
	}
}
