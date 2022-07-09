package main

import (
	"log"

	"github.com/serjbibox/rest-api-template/server"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	srv := server.New(viper.GetString("port"))
	srv.InitRoutes()
	srv.Run()
}

func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
}
