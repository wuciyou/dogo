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
)

type runLevel string

const (
	RUN_INFO    runLevel = "INFO"
	RUN_WARNING          = "WARNING"
	RUN_DEBUG            = "DEBUG"
	RUN_ERROR            = "ERROR"
)

var (
	RUN_INFO_FORMAT    = fmt.Sprintf("%c[0,0,%dm [INFO] %c[0m", 0x1B, 32, 0x1B)
	RUN_WARNING_FORMAT = fmt.Sprintf("%c[0,0,%dm [WARNING] %c[0m", 0x1B, 35, 0x1B)
	RUN_DEBUG_FORMAT   = fmt.Sprintf("%c[0,0,%dm [DEBUG] %c[0m", 0x1B, 36, 0x1B)
	RUN_ERROR_FORMAT   = fmt.Sprintf("%c[0,0,%dm [ERROR] %c[0m", 0x1B, 31, 0x1B)
)
