package errutil

import (
	"runtime"
	"strings"
)

type ScrapeError struct {
	Op   string // Operation being performed
	Err  error  // Original error
	File string // Source file where error occurred
	Line int    // Line number where error occurred
}

func (se *ScrapeError) Error() string {
	parts := []string{se.Op}
	
	if se.Err != nil {
		parts = append(parts, "error: "+se.Err.Error())
	}

	return strings.Join(parts, " - ")
}

func WrapError(op string, err error) error {
	if err == nil {
		return nil
	}

	_, file, line, _ := runtime.Caller(1)

	serr := &ScrapeError{
		Op:   op,
		Err:  err,
		File: file,
		Line: line,
	}

	return serr
}
