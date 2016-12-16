package dogo

type dogoConfig struct {
	Port             string
	ControllerSuffix string
	RunLevel         runLevel
}

type runLevel int

const (
	RUN_INFO runLevel = iota
	RUN_WARNING
	RUN_DEBUG
	RUN_ERROR
)

var RunTimeConfig dogoConfig

func (c dogoConfig) IsDebug() bool {
	return c.RunLevel == RUN_DEBUG
}

func init() {
	RunTimeConfig.RunLevel = RUN_DEBUG
	RunTimeConfig.Port = "8080"
	RunTimeConfig.ControllerSuffix = "Controller"
}
