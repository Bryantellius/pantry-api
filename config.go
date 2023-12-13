package main

import (
	"github.com/spf13/viper"
	"log"
)

// getConfig returns the configuration value for the given key
func getConfig(key string) string {
	return viper.GetString(key)
}

// SetupEnv loads and reads the configuration file
func SetupEnv() {
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}
