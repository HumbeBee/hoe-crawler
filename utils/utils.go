package utils

import (
	"regexp"
	"strings"
)

func GetHoeIDFromUrl(url string) string {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	re := regexp.MustCompile(`/(\d+)/`)
	match := re.FindStringSubmatch(url)
	if len(match) >= 2 {
		return match[1]
	}
	return ""
}
