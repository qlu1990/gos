package gos

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//Level log level
type Level int

var (
	// log file
	f *os.File
	//DefaultPrefix default
	DefaultPrefix = ""
	//DefaultCallerDepth 2
	DefaultCallerDepth = 2

	logger     *log.Logger
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

//Level is int
// start 0
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

//Log log struct type
type Log struct {
	logger    *log.Logger
	filePath  string
	fileName  string
	logPrefix string
	f         *os.File
}

//Glog default Glog var
var Glog = &Log{
	logger:    log.New(os.Stdout, DefaultPrefix, log.LstdFlags),
	logPrefix: "",
}

//Mlog default log middeware
var Mlog = Middleware{
	Name: "log",
	HandlerFunc: func(c *Context) {
		Glog.Info("Request : ", c.Request.Method, " ", c.Request.RequestURI)
	},
}

//LogSetOutPut set up log output
func (l *Log) LogSetOutPut(logFile string) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		Glog.Error("fail open logfile: ", logFile, "error: ", err)
		return
	}
	l.f = f
	l.logger.SetOutput(l.f)

}

//Close close the log file
func (l *Log) Close() {
	if l.f != nil {
		l.Close()
	}
}

//Debug Debug
func (l *Log) Debug(v ...interface{}) {
	l.setPrefix(DEBUG)
	l.logger.Println(v...)
}

//Info Info
func (l *Log) Info(v ...interface{}) {
	l.setPrefix(INFO)
	l.logger.Println(v...)
}

//Warn Warn
func (l *Log) Warn(v ...interface{}) {
	l.setPrefix(WARNING)
	l.logger.Println(v...)
}

//Error Error
func (l *Log) Error(v ...interface{}) {
	l.setPrefix(ERROR)
	l.logger.Println(v...)
}

//Fatal Fatal
func (l *Log) Fatal(v ...interface{}) {
	l.setPrefix(FATAL)
	l.logger.Fatalln(v...)
}

//Debug Debug
func Debug(v ...interface{}) {
	Glog.Debug(v...)
}

//Info Info
func Info(v ...interface{}) {
	Glog.Info(v...)
}

//Warn Warn
func Warn(v ...interface{}) {
	Glog.Warn(v...)
}

//Error Error
func Error(v ...interface{}) {
	Glog.Error(v...)
}

//Fatal Fatal
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
