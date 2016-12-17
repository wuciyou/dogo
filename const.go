package dogo

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST              = "POST"
	PUT               = "PUT"
	DELETE            = "DELETE"
)

type PipelineKey string

const (
	// 日志记录
	PIPELINE_LOG PipelineKey = "PIPELINE_LOG"
	// 路由解析
	PIPELINE_CONTEXT = "PIPELINE_CONTEXT"
)
