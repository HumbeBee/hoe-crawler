package main

import (
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
	"time"
)

// This cmd sets a long timeout (30 mins) since ConnectToPage will also handle browser download if needed
func main() {
	browser.ConnectToPage("https://google.com.vn", 30*time.Minute)
}
