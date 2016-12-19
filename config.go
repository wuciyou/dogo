package dogo

type dogoConfig struct {
	// 运行端口
	Port string
	// 运行级别
	RunLevel runLevel
	// 使用session
	UserSession bool
	// session 名称
	SessionName string
	// log 日志缓存大小(谨慎操作。默认值为0，如果 > 0 ，日志记录将是无序的，不能保证写入顺序)
	LogDataChanSize int
}

var RunTimeConfig dogoConfig

func (c dogoConfig) IsDebug() bool {
	return c.RunLevel == RUN_DEBUG
}

func init() {
	RunTimeConfig.RunLevel = RUN_DEBUG
	RunTimeConfig.Port = "8080"
	// 默认开启Session
	RunTimeConfig.UserSession = true
	// session 名称
	RunTimeConfig.SessionName = "DogoSessionID"
	//
	RunTimeConfig.LogDataChanSize = 0

}

// I don’t understand the start means
