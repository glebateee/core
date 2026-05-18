package main

import (
	"github.com/glebateee/core/logging"
)

func writeMessage(logger logging.Logger, msg string) {
	logger.Info(msg)

}

func main() {
	logger := logging.NewDefaultLogger(logging.Information)
	writeMessage(logger, "Hello")
}
