package main

import (
	"github.com/golobby/container/v2"
	"golang-container-to-wire.com/pkg/adapter"
	"golang-container-to-wire.com/pkg/implements"
)

func main() {
	// log := Container()
	log := Wire()

	log.Info("Hello World")
}

func Container() implements.Logger {
	adapter.Register()

	var logger implements.Logger
	container.Bind(&logger)

	return logger
}

func Wire() implements.Logger {
	return adapter.CreateLogger()
}
