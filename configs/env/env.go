package env

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnv(name string, path string) *viper.Viper {
	env := viper.New()

	env.AddConfigPath(path)
	env.SetConfigName(name)
	env.SetConfigType("env")

	err := env.ReadInConfig()
	if err != nil {
		log.Panicf("Error loading config file: %s", err)
	}

	return env
}
