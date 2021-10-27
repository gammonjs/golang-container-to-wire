package viper

import (
	"fmt"

	adaptee "github.com/spf13/viper"
	"golang-container-to-wire.com/pkg/implements"
)

type Viper struct {
	implements.Environment
}

func (v Viper) ServerAddress() string {
	return fmt.Sprintf("%s", adaptee.Get("address"))
}

func (v Viper) ServerPort() string {
	return fmt.Sprintf("%s", adaptee.Get("port"))
}

func NewEnvironment() Viper {
	adaptee.SetConfigName("local")  // name of config file (without extension)
	adaptee.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	adaptee.AddConfigPath("../env") // path to look for the config file in
	adaptee.AddConfigPath("./env")  // path to look for the config file in
	adaptee.AddConfigPath(".")      // optionally look for config in the working directory
	err := adaptee.ReadInConfig()   // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		fmt.Println(err)
	}

	return Viper{}
}
