package main

import (
	"log"
	"os"
)

type LogLevel int

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger
	logLevel LogLevel //LogLevel is enum
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERR ", msg)
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")  //2022/10/10 23:04:14 WARN Hello
	logger.Errorln("World") //2022/10/10 23:04:14 ERR World
}
