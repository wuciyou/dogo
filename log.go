package dogo

import (
	"fmt"
	"io"
	"log"
	"os"
)

type dogoLog struct {
}

var DogoLog = &dogoLog{}

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

// These functions write to the standard logger.

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dogoLog) Print(v ...interface{}) {
	log.Output(2, fmt.Sprint(v...))
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dogoLog) Printf(format string, v ...interface{}) {
	log.Output(2, fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (l *dogoLog) Println(v ...interface{}) {
	log.Output(2, fmt.Sprintln(v...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func (l *dogoLog) Fatal(v ...interface{}) {
	log.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func (l *dogoLog) Fatalf(format string, v ...interface{}) {
	log.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func (l *dogoLog) Fatalln(v ...interface{}) {
	log.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func (l *dogoLog) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	log.Output(2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func (l *dogoLog) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	log.Output(2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func (l *dogoLog) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	log.Output(2, s)
	panic(s)
}
