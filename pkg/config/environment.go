package config

import (
	"fmt"
	"strings"
)

type Environment string

const (
	Local       Environment = "local"
	Dev         Environment = "dev"
	Prod        Environment = "prod"
	Integration Environment = "integration"
)

func ParseEnvironment(s string) Environment {
	switch strings.ToLower(s) {
	case "dev":
		return Dev
	case "prod":
		return Prod
	case "integration":
		return Integration
	case "local":
		return Local
	}
	panic(fmt.Sprintf("unknown environment: %s", s))
}

func (e Environment) Name() string {
	return string(e)
}
