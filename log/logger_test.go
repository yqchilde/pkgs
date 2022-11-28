package log

import (
	"testing"

	"github.com/yqchilde/pkgs/config"
)

func TestNewLogger(t *testing.T) {
	config.New(".")
	var logConf Config
	if err := config.Load("config", &logConf); err != nil {
		t.Fatal(err)
	}

	logger := Init(&logConf)
	for i := 0; i < 1000; i++ {
		logger.Println("test")
	}
}
