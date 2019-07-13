package gos

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Log struct {
	logger    *log.Logger
	filePath  string
	fileName  string
	logPrefix string
}

var Glog = &Log{
	logger:    log.New(os.Stdout, DefaultPrefix, log.LstdFlags),
	logPrefix: "",
}
var Mlog = Middleware{
	Name: "log",
	HandlerFunc: func(c *Context) {
		Glog.Info("Request : ", c.Request.Method, " ", c.Request.RequestURI)
	},
}

func (l *Log) LogSetOutPut(logFile string) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		Glog.Error("fail open logfile: ", logFile, "error: ", err)
		return
	}
	l.logger.SetOutput(f)

}
func (l *Log) Debug(v ...interface{}) {
	l.setPrefix(DEBUG)
	l.logger.Println(v...)
}

func (l *Log) Info(v ...interface{}) {
	l.setPrefix(INFO)
	l.logger.Println(v...)
}

func (l *Log) Warn(v ...interface{}) {
	l.setPrefix(WARNING)
	l.logger.Println(v...)
}

func (l *Log) Error(v ...interface{}) {
	l.setPrefix(ERROR)
	l.logger.Println(v...)
}

func (l *Log) Fatal(v ...interface{}) {
	l.setPrefix(FATAL)
	l.logger.Fatalln(v...)
}

func Debug(v ...interface{}) {
	Glog.Debug(v...)
}

func Info(v ...interface{}) {
	Glog.Info(v...)
}

func Warn(v ...interface{}) {
	Glog.Warn(v...)
}

func Error(v ...interface{}) {
	Glog.Error(v...)
}

func Fatal(v ...interface{}) {
	Glog.Fatal(v...)
}

func (l *Log) setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		l.logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		l.logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	l.logger.SetPrefix(l.logPrefix)
}
