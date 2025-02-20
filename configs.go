package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey  string
	BaseURL string
	Port    string
}

func loadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Print(err)
	}

	viper.SetDefault("BaseURL", "https://api.brawlstars.com/v1")
	viper.SetDefault("Port", "8080")

	return
}
