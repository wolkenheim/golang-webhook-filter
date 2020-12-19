package main

import (
	"github.com/spf13/viper"
	"fmt"
)

func readConfig(envName string){
	viper.SetConfigName(envName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	viper.SetDefault("server.port", ":3000")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}