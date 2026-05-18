package main

import (
	"github.com/glebateee/core/config"
	"github.com/glebateee/core/logging"
)

func writeMessage(logger logging.Logger, cfg config.Config) {
	section, found := cfg.GetSection("main")
	if !found {
		logger.Warn("main section not found")
	} else {
		msg, ok := section.GetString("message")
		if ok {
			logger.Debug(msg)
		} else {
			logger.Info("default message")
		}
	}
}

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}
	logger := logging.NewDefaultLogger(cfg)
	writeMessage(logger, cfg)
}
