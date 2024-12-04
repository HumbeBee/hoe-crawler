package utils

import (
	"encoding/json"
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

func FormatJSON(data interface{}) string {
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error formatting: %v", err)
	}

	return string(pretty)
}

func PrintJSON(data interface{}) {
	fmt.Println(FormatJSON(data))
}
