//go:build wireinject

package adapter

import (
	"github.com/google/wire"
	"golang-container-to-wire.com/pkg/adapter/logrus"
	"golang-container-to-wire.com/pkg/adapter/viper"
)

func CreateEnvironment() viper.Viper {
	wire.Build(viper.CreateAdapter)

	return viper.Viper{}
}

func CreateLogger() logrus.Logrus {
	wire.Build(
		viper.CreateAdapter,
		logrus.CreateAdapter,
	)

	return logrus.Logrus{}
}
