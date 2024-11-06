package config

import "log"

const (
	BaseUrl          = "https://www.gaito.love"
	RequestPerSecond = 1.0
)

// InitLogger configures the standard logger with:
// - Date and time (2006/01/02 15:04:05)
// - File name and line number (file.go:123)
// This helps with debugging by showing exactly where and when each log occurred
func InitLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
