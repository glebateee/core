package logging

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

type Logger interface {
	Trace(string)
	Tracef(string, ...any)

	Debug(string)
	Debugf(string, ...any)

	Info(string)
	Infof(string, ...any)

	Warn(string)
	Warnf(string, ...any)

	Panic(string)
	Panicf(string, ...any)
}
