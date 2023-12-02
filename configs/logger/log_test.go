package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {

	logger := NewLogger()

	assert.NotNil(t, logger)
}
