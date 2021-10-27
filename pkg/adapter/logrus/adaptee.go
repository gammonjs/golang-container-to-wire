package logrus

import (
	adaptee "github.com/sirupsen/logrus"
	"golang-container-to-wire.com/pkg/adapter/viper"
	"golang-container-to-wire.com/pkg/implements"
)

type Logrus struct {
	implements.Logger
	Environment implements.Environment
}

func (l Logrus) Info(args ...interface{}) {
	adaptee.Info(l.Environment.ServerAddress(), args)
}

func CreateAdapter(environment viper.Viper) Logrus {
	return Logrus{Environment: environment}
}
