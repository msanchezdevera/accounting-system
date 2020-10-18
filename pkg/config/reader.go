package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const environmentKey = "ENVIRONMENT"

func ParseConfig() *Configuration {
	envVar, exists := os.LookupEnv(environmentKey)
	if !exists {
		envVar = "local"
	}

	environment := ParseEnvironment(envVar)

	configFile := fmt.Sprintf("config/%s/config.json", environment)

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var conf Configuration
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return &conf
}
