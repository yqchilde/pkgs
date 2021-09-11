package env

import "os"

const (
	// DevEnv development environment
	DevEnv = "dev"
	// ProEnv production environment
	ProEnv = "pro"
	// TestEnv local test environment
	TestEnv = "test"
)

func GetEnv() string {
	env := os.Getenv("env_mode")
	if env == "" {
		return TestEnv
	}

	return env
}
