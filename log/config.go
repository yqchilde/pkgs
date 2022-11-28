package log

type Config struct {
	Encoding              string // 编码
	Level                 string // 级别
	Writers               string // 写入器
	LogName               string // 日志名字
	LogRollingPolicy      string // 滚动策略
	LogRollingTimePattern string // 滚动时间格式，crontab格式(6位)
	LogRollingVolumeSize  string // 滚动卷大小
	LogBackupCount        int    // 备份数量
}
