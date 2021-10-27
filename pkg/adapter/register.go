package adapter

import (
	"github.com/golobby/container/v2"
	"golang-container-to-wire.com/pkg/adapter/logrus"
	"golang-container-to-wire.com/pkg/adapter/viper"
	"golang-container-to-wire.com/pkg/implements"
)

func Register() {
	container.Transient(func() implements.Environment {
		return viper.CreateAdapter()
	})

	container.Transient(func() implements.Logger {
		logger := logrus.Logrus{}
		container.Bind(&logger.Environment)
		return logger
	})
}
