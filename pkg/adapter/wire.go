//go:build wireinject

package adapter

import (
	"github.com/google/wire"
	"golang-container-to-wire.com/pkg/adapter/viper"
)

func Init() viper.Viper {
	wire.Build(viper.NewEnvironment)
	return viper.Viper{}
}
