package logger

import (
	"github.com/rs/zerolog"
	"os"
)

// Logger struct to be used as the logger object
type Logger struct {
	Lgr zerolog.Logger
}

// CreateLogger - create a zerolog logger
func CreateLogger(debugLevel string) *Logger {

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debugLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	l := &Logger{
		Lgr: logger,
	}
	return l
}

// Info - print out an info-level message
func (l *Logger) Info(msg string) {
	l.Lgr.Info().Msg(msg)
}

// Warn - print out an warning-level message
func (l *Logger) Warn(msg string) {
	l.Lgr.Warn().Msg(msg)
}

// Debug - print out an debug-level message
func (l *Logger) Debug(msg string) {
	l.Lgr.Debug().Msg(msg)
}

// Error - print out an error-level message
func (l *Logger) Error(msg string) {
	l.Lgr.Error().Msg(msg)
}
