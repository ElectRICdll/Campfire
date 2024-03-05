package log

import (
	"log"
	"os"
)

var Log Logger = NewLogger(0)

type Logger interface {
	Debug(string)

	Debugf(string, ...any)

	Info(string)

	Infof(string, ...any)

	Warning(string)

	Warningf(string, ...any)

	Error(string)

	Errorf(string, ...any)

	Critical(string)

	Criticalf(string, ...any)
}

type logger struct {
	level int
}

func NewLogger(level int) Logger {
	return &logger{level: level}
}

func (l *logger) logMessage(level int, message string) {
	if level >= l.level {
		log.Println(message)
	}
}

func (l *logger) logMessagef(level int, format string, v ...any) {
	if level >= l.level {
		log.Printf(format, v...)
	}
}

func (l *logger) Debug(message string) {
	l.logMessage(DEBUG, "[DEBUG]"+message)
}

func (l *logger) Debugf(format string, v ...any) {
	l.logMessagef(DEBUG, "[DEBUG]"+format, v...)
}

func (l *logger) Info(message string) {
	l.logMessage(INFO, "[INFO]"+message)
}

func (l *logger) Infof(format string, v ...any) {
	l.logMessagef(INFO, "[INFO]"+format, v...)
}

func (l *logger) Warning(message string) {
	l.logMessage(WARNING, "[WARNING]"+message)
}

func (l *logger) Warningf(format string, v ...any) {
	l.logMessagef(WARNING, "[WARNING]"+format, v...)
}

func (l *logger) Error(message string) {
	l.logMessage(ERROR, "[ERROR]"+message)
}

func (l *logger) Errorf(format string, v ...any) {
	l.logMessagef(ERROR, "[ERROR]"+format, v...)
}

func (l *logger) Critical(message string) {
	l.logMessage(CRITICAL, "[Critical]"+message)
	os.Exit(1)
}

func (l *logger) Criticalf(format string, v ...any) {
	l.logMessagef(CRITICAL, "[Critical]"+format, v...)
	os.Exit(1)
}
