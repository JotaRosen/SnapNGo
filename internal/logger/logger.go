package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
)

// Wrapper for zerolog lib
type Logger struct {
	*zerolog.Logger
	component string
}

// NewLogger creates a new logger
func NewLogger(logFile, component string) (*Logger, error) {
	writers := []io.Writer{os.Stdout}

	if logFile != "" {
		if err := os.MkdirAll(filepath.Dir(logFile), 0755); err != nil {
			return nil, fmt.Errorf("unable to create log file directory: %w", err)
		}
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("unable to open log file: %w", err)
		}
		writers = append(writers, file)
	}

	multiWriter := io.MultiWriter(writers...)
	logger := zerolog.New(multiWriter).With().Timestamp().Logger()

	return &Logger{
		Logger:    &logger,
		component: component,
	}, nil
}

// WithComponent builds a logger with a specific component
func (l *Logger) WithComponent(component string) *Logger {
	l.component = component
	return l
}
