package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

//log.New(logFile, "[debug]", log.Ldate|log.Ltime|log.Llongfile)
// 第一个参数为输出io，可以是文件也可以是实现了该接口的对象，此处为日志文件；
// 第二个参数为自定义前缀；第三个参数为输出日志的格式选项，可多选组合
// 第三个参数可选如下：
/*
   Ldate         = 1             // 日期：2009/01/23
   Ltime         = 2             // 时间：01:23:23
   Lmicroseconds = 4             // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
   Llongfile     = 8             // 文件全路径名+行号： /a/b/c/d.go:23
   Lshortfile    = 16            // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
   LstdFlags     = Ldate | Ltime // 标准logger的初始值
*/

type LoggerI interface {
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})
	Infoln(v ...interface{})
	Infof(format string, v ...interface{})
	Warnln(v ...interface{})
	Warnf(format string, v ...interface{})
	Errorln(v ...interface{})
	Errorf(format string, v ...interface{})
	Fataln(v ...interface{})
	Fatalf(format string, v ...interface{})
}

const calldepth = 3

type LogLevel int

const (
	TRACE LogLevel = iota // TRACE 用户级基本输出
	DEBUG                 // DEBUG 用户级调试输出
	INFO                  // INFO 用户级重要信息
	WARN                  // WARN 用户级警告信息
	ERROR                 // ERROR 用户级错误信息
	FATAL
)

var LevelString = [...]string{
	"[TRACE] ",
	"[DEBUG] ",
	"[INFO]  ",
	"[WARN]  ",
	"[ERROR] ",
	"[FATAL] ",
}

type Logger struct {
	fPath  string
	logger *log.Logger
}

func NewLogger(basePath, fileName string) *Logger {
	return newLogger(basePath, fileName)
}

func newLogger(basePath, fileName string) *Logger {
	//var logFile io.Writer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()

	dir := fmt.Sprintf("%s/%04d-%02d-%02d", basePath, year, month, day)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fPath := fmt.Sprintf("%s/%s.%02d.%02d.%02d.log", dir, fileName, hour, min, sec)

	// 第三个参数为文件权限，请参考linux文件权限，664在这里为八进制，代表：rw-rw-r--
	logFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(logFile, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)

	return &Logger{
		fPath:  fPath,
		logger: logger,
	}
}

func header(lev LogLevel, msg string) string {
	prefix := LevelString[lev]
	return fmt.Sprintf("%s %s", prefix, msg)
}

func (l *Logger) print(lev LogLevel, format string, v ...interface{}) {

	if l.logger != nil {
		prefix := LevelString[lev]
		if l.logger.Prefix() != prefix {
			l.logger.SetPrefix(prefix)
		}

		text := ""
		if format == "" {
			text = fmt.Sprintln(v...)
		} else {
			text = fmt.Sprintf(format, v...)
		}

		l.logger.Output(calldepth, text)
	}

}

func (l *Logger) Debugln(v ...interface{}) {
	l.print(DEBUG, "", v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.print(DEBUG, format, v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.print(INFO, "", v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.print(INFO, format, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	l.print(WARN, "", v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.print(WARN, format, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.print(ERROR, "", v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.print(ERROR, format, v...)
}

func (l *Logger) Fataln(v ...interface{}) {
	l.print(FATAL, "", v...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.print(FATAL, format, v...)
	os.Exit(1)
}
