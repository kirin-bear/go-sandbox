package main

import (
	"log"
	"os"
)

type LogLevel int

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
	Poper
}

type Poper interface {
	Pop() string
}



func (l LogLevel) IsValid() bool {
    switch l {
    case LogLevelInfo, LogLevelWarning, LogLevelError:
        return true
    default:
        return false
    }
}

func (le *LogExtended) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
        return
    }
	le.logLevel = logLvl
}

func (le LogExtended) Infoln(msg string) {
	le.println(LogLevelInfo, "INFO ", msg)
}

func (le LogExtended) Warnln(msg string) {
	le.println(LogLevelWarning, "WARN ", msg)
}

func (le LogExtended) Errorln(msg string) {
	le.println(LogLevelError, "ERR ", msg)
}

func (le *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
    // игнорируем сообщения, если уровень логгера меньше scrLogLvl
    // ...
	if (le.logLevel < srcLogLvl) {
		return
	}

    le.Logger.Println(prefix + msg)
}

func NewLogExtended() *LogExtended {
    return &LogExtended{
        Logger:   log.New(os.Stderr, "", log.LstdFlags),
        logLevel: LogLevelError,
    }
}

const (
    LogLevelError LogLevel = iota
    LogLevelWarning
    LogLevelInfo
) 

func main() {

	nle := NewLogExtended()
    nle.SetLogLevel(LogLevelWarning)
    nle.Infoln("Не должно напечататься")
    nle.Warnln("Hello")
    nle.Errorln("World")
    nle.Println("Debug")
}