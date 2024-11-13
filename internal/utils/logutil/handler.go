package logutil

import (
	"log"
	"os"
)

type Logger struct {
	logger   *log.Logger
	minLevel LogLevel
}

func NewLogger(minLevel LogLevel) *Logger {
	return &Logger{
		logger:   log.Default(),
		minLevel: minLevel,
	}
}

// Internal logging method
func (l *Logger) Log(level LogLevel, msg string, args ...interface{}) {
	if level >= l.minLevel {
		l.logger.Printf("["+level.String()+"] "+msg, args...)
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Log(DEBUG, msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Log(INFO, msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Log(WARN, msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Log(ERROR, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Log(FATAL, msg, args...)
	// panic(msg) // or os.Exit(1) depending on your needs
	os.Exit(1)
}
