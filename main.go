package main

import (
	"fmt"

	"github.com/glebateee/core/config"
	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/placeholder"
	"github.com/glebateee/core/services"
)

func writeMessage(
	prefix string,
	logger logging.Logger,
	cfg config.Config,
) {
	fmt.Println(prefix)
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
	services.RegisterDefaultServices()
	placeholder.Start()
}
