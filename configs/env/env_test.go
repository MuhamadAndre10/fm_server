package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("LoadEnv() panic = %v", r)
		}
	}()

	// Load env dari file config.env
	env := LoadEnv("config", "../../")

	assert.NotNil(t, env)
}
