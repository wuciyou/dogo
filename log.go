package dogo

import (
	"fmt"
	"io"
	"log"
	"runtime/debug"
)

type RunLevel string

var (
	RUN_INFO    RunLevel = "INFO"
	RUN_WARNING RunLevel = "WARN"
	RUN_DEBUG   RunLevel = "DEBUG"
	RUN_ERROR   RunLevel = "ERROR"

	RUN_INFO_FORMAT    = fmt.Sprintf("%c[0,0,%dm %-7s", 0x1B, 32, RUN_INFO)
	RUN_WARNING_FORMAT = fmt.Sprintf("%c[0,0,%dm %-7s", 0x1B, 35, RUN_WARNING)
	RUN_DEBUG_FORMAT   = fmt.Sprintf("%c[0,0,%dm %-7s", 0x1B, 36, RUN_DEBUG)
	RUN_ERROR_FORMAT   = fmt.Sprintf("%c[0,0,%dm %-7s", 0x1B, 31, RUN_ERROR)
	_CLEAR_COLOR       = fmt.Sprintf("%c[0m", 0x1B)
)

type dglog struct {
	logDataChanSize int
	runLevel        RunLevel
}

var Dglog = &dglog{logDataChanSize: 0, runLevel: RUN_DEBUG}

func init() {
	SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func (l *dglog) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func (l *dglog) Output(s string) {
	log.Output(3, s)
}

func (l *dglog) Flags() int {
	return log.Flags()
}

func (l *dglog) SetFlags(flag int) {
	log.SetFlags(flag)
}

func (l *dglog) Prefix() string {
	return log.Prefix()
}

func (l *dglog) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

func (l *dglog) Info(v ...interface{}) {
	l.SetPrefix(RUN_INFO_FORMAT)
	v = append(v, _CLEAR_COLOR)
	l.Output(fmt.Sprint(v...))
}

func (l *dglog) Infof(format string, v ...interface{}) {
	l.SetPrefix(RUN_INFO_FORMAT)
	l.Output(fmt.Sprintf(format+_CLEAR_COLOR, v...))
}

func (l *dglog) Debug(v ...interface{}) {

	if l.runLevel == RUN_DEBUG {
		l.SetPrefix(RUN_DEBUG_FORMAT)
		l.Output(fmt.Sprint(v...) + _CLEAR_COLOR)
	}
}

func (l *dglog) Debugf(format string, v ...interface{}) {

	if l.runLevel == RUN_DEBUG {
		l.SetPrefix(RUN_DEBUG_FORMAT)
		l.Output(fmt.Sprintf(format, v...) + _CLEAR_COLOR)
	}
}

func (l *dglog) Warning(v ...interface{}) {
	if l.runLevel == RUN_DEBUG {
		l.SetPrefix(RUN_WARNING_FORMAT)
		l.Output(fmt.Sprint(v...) + _CLEAR_COLOR)
	}
}

func (l *dglog) Warningf(format string, v ...interface{}) {
	if l.runLevel == RUN_DEBUG {
		l.SetPrefix(RUN_WARNING_FORMAT)
		l.Output(fmt.Sprintf(format, v...) + _CLEAR_COLOR)
	}
}

func (l *dglog) Error(v ...interface{}) {
	l.SetPrefix(RUN_ERROR_FORMAT)
	s := fmt.Sprint(v...) + string(debug.Stack()) + _CLEAR_COLOR
	l.Output(s)
}

func (l *dglog) Errorf(format string, v ...interface{}) {
	l.SetPrefix(RUN_ERROR_FORMAT)
	s := fmt.Sprintf(format+string(debug.Stack())+_CLEAR_COLOR, v...)

	l.Output(s)

}

func SetOutput(w io.Writer) {
	Dglog.SetOutput(w)
}

func Flags() int {
	return Dglog.Flags()
}

func SetFlags(flag int) {
	Dglog.SetFlags(flag)
}

func Prefix() string {
	return Dglog.Prefix()
}

func SetPrefix(prefix string) {
	Dglog.SetPrefix(prefix)
}

func Info(v ...interface{}) {
	Dglog.Info(v...)
}

func Infof(format string, v ...interface{}) {
	Dglog.Infof(format, v...)
}

func Debug(v ...interface{}) {
	Dglog.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	Dglog.Debugf(format, v...)
}

func Warning(v ...interface{}) {
	Dglog.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	Dglog.Warningf(format, v...)
}

func Error(v ...interface{}) {
	Dglog.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	Dglog.Errorf(format, v...)
}
