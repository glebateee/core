package logging

import (
	"log"
	"os"
	"strings"

	"github.com/glebateee/core/config"
)

func NewDefaultLogger(cfg config.Config) *defaultLogger {
	level := cfg.GetStringDefault("logging:level", "debug")
	flags := log.Lmsgprefix | log.Ltime
	return &defaultLogger{
		minLevel: LogLevelFromString(level),
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, "TRACE ", flags),
			Debug:       log.New(os.Stdout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning:     log.New(os.Stdout, "WARN ", flags),
			Fatal:       log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}

func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "information":
		level = Information
	case "warning":
		level = Warning
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}
