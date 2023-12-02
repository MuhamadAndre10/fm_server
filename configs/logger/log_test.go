package logger

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewLogger(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in f", r)
		}
	}()

	logger := NewLogger("../../../log.txt")

	assert.NotNil(t, logger)
}
