package env

import "os"

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

var (
	env = os.Getenv("SERVER_ENV")
)

func GetEnv() string {
	if env == "" {
		env = Development
	}
	return env
}

func SetEnv(environment string) {
	env = environment
}
