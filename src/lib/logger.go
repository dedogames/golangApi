package lib

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

type LoggerObj struct {
	level  LogLevel
	writer io.Writer
}

func NewLogger(level LogLevel, writer io.Writer) *LoggerObj {
	return &LoggerObj{
		level:  level,
		writer: writer,
	}
}

func (l *LoggerObj) Log(level LogLevel, message string) {
	if level >= l.level {
		levelType := "DEBUB"
		switch level {
		case Info:
			levelType = "INFO"
		case Warn:
			levelType = "WARN"
		case Error:
			levelType = "ERROR"
		}
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logString := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelType, message)
		l.writer.Write([]byte(logString))
	}
}

func (l *LoggerObj) Debug(message string) {
	l.Log(Debug, message)
}

func (l *LoggerObj) Info(message string) {
	l.Log(Info, message)
}

func (l *LoggerObj) Warn(message string) {
	l.Log(Warn, message)
}

func (l *LoggerObj) Error(message string) {
	l.Log(Error, message)
}

var Logger *LoggerObj = NewLogger(Info, os.Stdout)
