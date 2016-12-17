package dogo

import (
	"fmt"
	"io"
	"log"
)

/**
*
   // 前景 背景 颜色
   // ---------------------------------------
   // 30  40  黑色
   // 31  41  红色
   // 32  42  绿色
   // 33  43  黄色
   // 34  44  蓝色
   // 35  45  紫红色
   // 36  46  青蓝色
   // 37  47  白色
   //
   // 代码 意义
   // -------------------------
   //  0  终端默认设置
   //  1  高亮显示
   //  4  使用下划线
   //  5  闪烁
   //  7  反白显示
   //  8  不可见

   for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
       for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
           for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
               fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
           }
           fmt.Println("")
       }
       fmt.Println("")
   }
*/

type dogoLog struct {
}

var DogoLog = &dogoLog{}

func init() {

	// Ldate         = 1 << iota     // the date: 2009/01/23
	// Ltime                         // the time: 01:23:23
	// Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	// Llongfile                     // full file name and line number: /a/b/c/d.go:23
	// Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	// LstdFlags

	// DogoLog.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	DogoLog.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
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
	log.Output(2, fmt.Sprint(v...))
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dogoLog) Infof(format string, v ...interface{}) {
	log.Output(2, fmt.Sprintf(RUN_INFO_FORMAT+format, v...))
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dogoLog) Debug(v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		v = append([]interface{}{RUN_DEBUG_FORMAT}, v...)
		log.Output(2, fmt.Sprint(v...))
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dogoLog) Debugf(format string, v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		log.Output(2, fmt.Sprintf(RUN_DEBUG_FORMAT+format, v...))
	}
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (l *dogoLog) Warning(v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		v = append([]interface{}{RUN_WARNING_FORMAT}, v...)
		log.Output(2, fmt.Sprint(v...))
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *dogoLog) Warningf(format string, v ...interface{}) {
	if RunTimeConfig.IsDebug() {
		log.Output(2, fmt.Sprintf(RUN_WARNING_FORMAT+format, v...))
	}
}

// Panic is equivalent to Print() followed by a call to panic().
func (l *dogoLog) Error(v ...interface{}) {
	v = append([]interface{}{RUN_ERROR_FORMAT, fmt.Sprintf("%c[0,0,%dm ", 0x1B, 31)}, v...)
	v = append(v, fmt.Sprintf("%c[0m", 0x1B))
	s := fmt.Sprint(v...)
	log.Output(2, s)

	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func (l *dogoLog) Errorf(format string, v ...interface{}) {
	v = append([]interface{}{0x1B, 31}, v...)
	v = append(v, 0x1B)
	s := fmt.Sprintf(RUN_ERROR_FORMAT+"%c[0,0,%dm "+format+"%c[0m ", v...)
	log.Output(2, s)
	panic(s)
}
