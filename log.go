package dogo

import (
	"fmt"
	"io"
	"log"
)

type dogoLog struct {
}

var DogoLog = &dogoLog{}

var logDataChan chan string

func init() {

	// DogoLog.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	DogoLog.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	logDataChan = make(chan string, RunTimeConfig.LogDataChanSize)
	go func() {
		for {
			log.Output(2, <-logDataChan)
		}

	}()
}

// SetOutput sets the output destination for the standard logger.
func (l *dogoLog) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
func (l *dogoLog) Flags() int {
	return log.Flags()
}

// SetFlags sets the output flags for the standard logger.
func (l *dogoLog) SetFlags(flag int) {
	log.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func (l *dogoLog) Prefix() string {
	return log.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func (l *dogoLog) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dogoLog) Info(v ...interface{}) {
	v = append([]interface{}{RUN_INFO_FORMAT}, v...)
	logDataChan <- fmt.Sprint(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dogoLog) Infof(format string, v ...interface{}) {
	logDataChan <- fmt.Sprintf(RUN_INFO_FORMAT+format, v...)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dogoLog) Debug(v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		v = append([]interface{}{RUN_DEBUG_FORMAT}, v...)
		logDataChan <- fmt.Sprint(v...)
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dogoLog) Debugf(format string, v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		logDataChan <- fmt.Sprintf(RUN_DEBUG_FORMAT+format, v...)
	}
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dogoLog) Warning(v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		v = append([]interface{}{RUN_WARNING_FORMAT}, v...)
		logDataChan <- fmt.Sprint(v...)
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dogoLog) Warningf(format string, v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		logDataChan <- fmt.Sprintf(RUN_WARNING_FORMAT+format, v...)
	}
}

// Panic is equivalent to Print() followed by a call to panic().
func (l *dogoLog) Error(v ...interface{}) {
	v = append([]interface{}{RUN_ERROR_FORMAT, fmt.Sprintf("%c[0,0,%dm ", 0x1B, 31)}, v...)
	v = append(v, fmt.Sprintf("%c[0m", 0x1B))
	s := fmt.Sprint(v...)
	logDataChan <- s
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func (l *dogoLog) Errorf(format string, v ...interface{}) {
	v = append([]interface{}{0x1B, 31}, v...)
	v = append(v, 0x1B)
	s := fmt.Sprintf(RUN_ERROR_FORMAT+"%c[0,0,%dm "+format+"%c[0m ", v...)
	logDataChan <- s
	panic(s)
}
