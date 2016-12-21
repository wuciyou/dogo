package config

type dogoConfig struct {
	// 运行端口
	// port int
	// 运行级别
	// 使用session
	UserSession bool
	// session 名称
	SessionName string
	// log 日志缓存大小(谨慎操作。默认值为0，如果 > 0 ，日志记录将是无序的，不能保证写入顺序)
	// 这里的缓存不是日志记录数据大小
	LogDataChanSize int
	// 默认 ajax 返回的格式
	// 支持 json, xml
	// 如果为空，将根据 请求路径后缀自动匹配
	ajaxReturnRormat string
	// web服务器名称
	// 默认为 DoGoServerv1
	serverName string
	// 静态资源请求路径
	// 默认为 /imgages,/css
	// 多个请求路径 使用,分割
	// 最终的保存路径为
	staticRequstPath string
	// 静态资源存放路径   staticRootPath + staticRequstPath
	// 默认为 项目根下面的 static目录
	staticRootPath string
}

var RunTimeConfig dogoConfig

func (c dogoConfig) AjaxReturnRormat() string {
	return c.ajaxReturnRormat
}
func (c dogoConfig) StaticRequstPath() string {
	return c.staticRequstPath
}
func (c dogoConfig) StaticRootPath() string {
	return c.staticRootPath
}

func (c dogoConfig) ServerName() string {
	return c.serverName
}

func init() {
	// 默认开启Session
	RunTimeConfig.UserSession = true
	// session 名称
	RunTimeConfig.SessionName = "DogoSessionID"
	// log 日志缓存队列大小
	RunTimeConfig.LogDataChanSize = 0
	// 默认 ajax 返回的格式
	RunTimeConfig.ajaxReturnRormat = "xml"
	RunTimeConfig.serverName = "DoGoServerv1"

	RunTimeConfig.staticRequstPath = "/static/"
	RunTimeConfig.staticRootPath = "./"
}

// I don’t understand the start means
