package log

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const (
	RollingTimeDaily  = "daily"
	RollingTimeHourly = "hourly"
)

// Write log by time
func getLogWriterWithTime(cfg *Config, filename string) io.Writer {
	logFullPath := filename
	rotationPolicy := cfg.LogRollingPolicy
	backupCount := cfg.LogBackupCount

	rollingDuration := time.Hour * 24
	if rotationPolicy == RollingTimeHourly {
		rollingDuration = time.Hour
	}

	hook, err := rotatelogs.New(
		logFullPath+".%Y%m%d%H",
		rotatelogs.WithLinkName(logFullPath),
		rotatelogs.WithRotationCount(backupCount),
		rotatelogs.WithRotationTime(rollingDuration),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
