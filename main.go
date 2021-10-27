package main

import (
	"fmt"

	"github.com/golobby/container/v2"
	"golang-container-to-wire.com/pkg/adapter"
	"golang-container-to-wire.com/pkg/implements"
)

func main() {
	adapter.Register()

	var environment implements.Environment
	container.Bind(&environment)

	// var environment implements.Environment
	// environment = adapter.Init()

	fmt.Println(environment.ServerAddress())
}
