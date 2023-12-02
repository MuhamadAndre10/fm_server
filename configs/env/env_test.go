package env

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadEnv(t *testing.T) {

	// Load env dari file config.env
	env := LoadEnv("config", "../../")

	assert.NotNil(t, env)
	assert.Equal(t, "FAVAA_CUSTOMER_APP", env.GetString("APP_NAME"))
}
