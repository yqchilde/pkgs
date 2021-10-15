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
	rollingPolicy := cfg.LogRollingPolicy
	backupCount := cfg.LogBackupCount

	var rollingDuration time.Duration
	if rollingPolicy == RollingTimeHourly {
		rollingDuration = time.Hour
	} else if rollingPolicy == RollingTimeDaily {
		rollingDuration = time.Hour * 24
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
