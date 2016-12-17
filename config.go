package dogo

type dogoConfig struct {
	// 运行端口
	Port string
	// 运行级别
	RunLevel runLevel
}

var RunTimeConfig dogoConfig

func (c dogoConfig) IsDebug() bool {
	return c.RunLevel == RUN_DEBUG
}

func init() {
	RunTimeConfig.RunLevel = RUN_DEBUG
	RunTimeConfig.Port = "8080"
}
