package env

import "os"

const (
	// EnvDev development environment
	EnvDev = "dev"
	// EnvPro production environment
	EnvPro = "pro"
	// EnvTest local test environment
	EnvTest = "test"
)

func GetEnv() string {
	env := os.Getenv("env_mode")
	if env == "" {
		return EnvTest
	}

	return env
}
