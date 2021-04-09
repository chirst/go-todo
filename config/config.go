package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// InitConfig reads in config from a file or panics on failure
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error in config: %s\n", err))
	}
}

// GetAddress returns a network address to listen for requests
func GetAddress() string {
	return viper.GetString("server_address")
}

// GetSignKey returns a secret key to sign JSON Web Tokens
func GetSignKey() string {
	return viper.GetString("jwt_sign_key")
}
