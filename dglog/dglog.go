package dglog

import (
	"fmt"
	"github.com/wuciyou/dogo/common"
	"io"
	"log"
)

var logDataChan chan string

type dglog struct {
	logDataChanSize int
	runLevel        common.RunLevel
}

var dglogObj = &dglog{logDataChanSize: 0, runLevel: common.RUN_DEBUG}

func init() {

	// DogoLog.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	logDataChan = make(chan string, dglogObj.logDataChanSize)
	go func() {
		for {
			log.Output(2, <-logDataChan)
		}
	}()
}

// SetOutput sets the output destination for the standard logger.
func (l *dglog) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
func (l *dglog) Flags() int {
	return log.Flags()
}

// SetFlags sets the output flags for the standard logger.
func (l *dglog) SetFlags(flag int) {
	log.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func (l *dglog) Prefix() string {
	return log.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func (l *dglog) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dglog) Info(v ...interface{}) {
	v = append([]interface{}{common.RUN_INFO_FORMAT}, v...)
	logDataChan <- fmt.Sprint(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dglog) Infof(format string, v ...interface{}) {
	logDataChan <- fmt.Sprintf(common.RUN_INFO_FORMAT+format, v...)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dglog) Debug(v ...interface{}) {
	if l.runLevel == common.RUN_DEBUG {
		v = append([]interface{}{common.RUN_DEBUG_FORMAT}, v...)
		logDataChan <- fmt.Sprint(v...)
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dglog) Debugf(format string, v ...interface{}) {
	if l.runLevel == common.RUN_DEBUG {
		logDataChan <- fmt.Sprintf(common.RUN_DEBUG_FORMAT+format, v...)
	}
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dglog) Warning(v ...interface{}) {
	if l.runLevel == common.RUN_DEBUG {
		v = append([]interface{}{common.RUN_WARNING_FORMAT}, v...)
		logDataChan <- fmt.Sprint(v...)
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dglog) Warningf(format string, v ...interface{}) {
	if l.runLevel == common.RUN_DEBUG {
		logDataChan <- fmt.Sprintf(common.RUN_WARNING_FORMAT+format, v...)
	}
}

// Panic is equivalent to Print() followed by a call to panic().
func (l *dglog) Error(v ...interface{}) {
	v = append([]interface{}{common.RUN_ERROR_FORMAT, fmt.Sprintf("%c[0,0,%dm ", 0x1B, 31)}, v...)
	v = append(v, fmt.Sprintf("%c[0m", 0x1B))
	s := fmt.Sprint(v...)
	logDataChan <- s
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func (l *dglog) Errorf(format string, v ...interface{}) {
	v = append([]interface{}{0x1B, 31}, v...)

	v = append(v, 0X1B)
	s := fmt.Sprintf(common.RUN_ERROR_FORMAT+"%c[0,0,%dm "+format+" %c[0m ", v...)
	logDataChan <- s
	panic(s)
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	dglogObj.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
func Flags() int {
	return dglogObj.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	dglogObj.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return dglogObj.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	dglogObj.SetPrefix(prefix)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	dglogObj.Info(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	dglogObj.Infof(format, v...)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	dglogObj.Debug(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	dglogObj.Debugf(format, v...)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Warning(v ...interface{}) {
	dglogObj.Warning(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, v ...interface{}) {
	dglogObj.Warningf(format, v...)
}

// Panic is equivalent to Print() followed by a call to panic().
func Error(v ...interface{}) {
	dglogObj.Error(v...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Errorf(format string, v ...interface{}) {
	dglogObj.Errorf(format, v...)
}
