package log

import (
	"errors"
	"io"

	"github.com/arthurkiller/rollingwriter"
)

const (
	// RollingPolicyTime 按时间滚动
	RollingPolicyTime = "time"

	// RollingPolicyVolume 按卷大小滚动
	RollingPolicyVolume = "volume"
)

func getLogRollingPolicy(conf *Config) (io.Writer, error) {
	rollingConf := &rollingwriter.Config{
		LogPath:           "./",
		FileName:          conf.LogName,
		TimeTagFormat:     "20060102150405",
		WriterMode:        "lock",
		MaxRemain:         conf.LogBackupCount,
		FilterEmptyBackup: false,
		Compress:          false,
	}

	switch conf.LogRollingPolicy {
	case RollingPolicyTime:
		rollingConf.RollingPolicy = rollingwriter.TimeRolling
		rollingConf.RollingTimePattern = conf.LogRollingTimePattern
	case RollingPolicyVolume:
		rollingConf.RollingPolicy = rollingwriter.VolumeRolling
		rollingConf.RollingVolumeSize = conf.LogRollingVolumeSize
	default:
		return nil, errors.New("unknown log rolling policy")
	}

	return rollingwriter.NewWriterFromConfig(rollingConf)
}
