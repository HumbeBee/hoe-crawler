package errutil

import (
	"log"
	"os"
	"runtime"
	"strings"
)

type ScrapeError struct {
	Op     string // Operation being performed
	Target string // Target being scraped (URL, selector, etc)
	Err    error  // Original error
	File   string // Source file where error occurred
	Line   int    // Line number where error occurred
}

type ErrorHandler struct {
	logger   *log.Logger
	minLevel LogLevel
}

func (se *ScrapeError) Error() string {
	parts := []string{se.Op}

	if se.Target != "" {
		parts = append(parts, "target: "+se.Target)
	}
	if se.Err != nil {
		parts = append(parts, "error: "+se.Err.Error())
	}

	return strings.Join(parts, " - ")
}

func NewErrorHandler(logger *log.Logger, minLevel LogLevel) *ErrorHandler {
	if logger == nil {
		logger = log.Default()
	}

	return &ErrorHandler{logger, minLevel}
}

func (h *ErrorHandler) WrapError(op string, err error, target string) error {
	if err == nil {
		return nil
	}

	_, file, line, _ := runtime.Caller(1)

	serr := &ScrapeError{
		Op:     op,
		Target: target,
		Err:    err,
		File:   file,
		Line:   line,
	}

	// h.log(ERROR, "%s at %s:%d - %v", op, file, line, err)

	return serr
}

func (h *ErrorHandler) Debug(msg string, args ...interface{}) {
	h.log(DEBUG, msg, args...)
}

func (h *ErrorHandler) Info(msg string, args ...interface{}) {
	h.log(INFO, msg, args...)
}

func (h *ErrorHandler) Warn(msg string, args ...interface{}) {
	h.log(WARN, msg, args...)
}

func (h *ErrorHandler) Error(msg string, args ...interface{}) {
	h.log(ERROR, msg, args...)
}

func (h *ErrorHandler) Fatal(msg string, args ...interface{}) {
	h.log(FATAL, msg, args...)
	// panic(msg) // or os.Exit(1) depending on your needs
	os.Exit(1)
}

// Internal logging method
func (h *ErrorHandler) log(level LogLevel, msg string, args ...interface{}) {
	if level >= h.minLevel {
		h.logger.Printf("["+level.String()+"] "+msg, args...)
	}
}
