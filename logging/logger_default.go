package logging

import (
	"fmt"
	"log"
)

type defaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *defaultLogger) write(level LogLevel, msg string) {
	if level >= l.minLevel {
		l.loggers[level].Output(2, msg)
	}
}

func (l *defaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

func (l *defaultLogger) Tracef(msg string, args ...any) {
	l.write(Trace, fmt.Sprintf(msg, args...))
}

func (l *defaultLogger) Debug(msg string) {
	l.write(Debug, msg)

}
func (l *defaultLogger) Debugf(msg string, args ...any) {
	l.write(Debug, fmt.Sprintf(msg, args...))

}

func (l *defaultLogger) Info(msg string) {
	l.write(Information, msg)

}
func (l *defaultLogger) Infof(msg string, args ...any) {
	l.write(Information, fmt.Sprintf(msg, args...))

}

func (l *defaultLogger) Warn(msg string) {
	l.write(Warning, msg)

}
func (l *defaultLogger) Warnf(msg string, args ...any) {
	l.write(Warning, fmt.Sprintf(msg, args...))

}

func (l *defaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

func (l *defaultLogger) Panicf(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}
