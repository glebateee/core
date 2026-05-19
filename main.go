package main

import (
	"github.com/glebateee/core/placeholder"
	"github.com/glebateee/core/services"
)

func main() {
	services.RegisterDefaultServices()
	placeholder.Start()
}
