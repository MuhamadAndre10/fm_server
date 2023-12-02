package env

import (
	"github.com/spf13/viper"
)

func LoadEnv(name string, path string) *viper.Viper {
	env := viper.New()

	env.AddConfigPath(path)
	env.SetConfigName(name)
	env.SetConfigType("env")

	err := env.ReadInConfig()
	if err != nil {
		panic("Failed to load configuration")
	}

	return env
}
