// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package adapter

import (
	"golang-container-to-wire.com/pkg/adapter/viper"
)

// Injectors from wire.go:

func Init() viper.Viper {
	viperViper := viper.NewEnvironment()
	return viperViper
}