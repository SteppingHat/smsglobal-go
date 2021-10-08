package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLogger(t *testing.T) {
	log := CreateLogger("debug")
	assert.NotNil(t, log, "Logger has been loaded")
}

// These tests are here for the sake of completeness and to raise the coverage as
// the functions are merely calling the functions of the third-party library
func TestLogging(t *testing.T) {

	log := CreateLogger("debug")
	assert.NotNil(t, log, "Logger has been loaded")

	log.Info("Logging Info")
	log.Warn("Logging Warn")
	log.Debug("Logging Debug")
	log.Error("Logging Error")
}
