package dogo

import (
	"fmt"
)

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
	// session
	PIPELINE_SESSION = "PIPELINE_SESSION"
)

type runLevel string

const (
	RUN_INFO    runLevel = "INFO"
	RUN_WARNING          = "WARNING"
	RUN_DEBUG            = "DEBUG"
	RUN_ERROR            = "ERROR"
)

var (
	RUN_INFO_FORMAT    = fmt.Sprintf("%c[0,0,%dm [%-7s] %c[0m", 0x1B, 32, RUN_INFO, 0x1B)
	RUN_WARNING_FORMAT = fmt.Sprintf("%c[0,0,%dm [%-7s] %c[0m", 0x1B, 35, RUN_WARNING, 0x1B)
	RUN_DEBUG_FORMAT   = fmt.Sprintf("%c[0,0,%dm [%-7s] %c[0m", 0x1B, 36, RUN_DEBUG, 0x1B)
	RUN_ERROR_FORMAT   = fmt.Sprintf("%c[0,0,%dm [%-7s] %c[0m", 0x1B, 31, RUN_ERROR, 0x1B)
)
