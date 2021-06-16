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
		panic(fmt.Errorf("fatal error in config: %s", err))
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

// GetPostgresSourceName returns a string containing connection info
// specific to a Postgres database
func GetPostgresSourceName() string {
	return "host=" + viper.GetString("pg_host") + " " +
		"port=" + viper.GetString("pg_port") + " " +
		"user=" + viper.GetString("pg_user") + " " +
		"password=" + viper.GetString("pg_password") + " " +
		"dbname=" + viper.GetString("pg_dbname") + " " +
		"sslmode=" + viper.GetString("pg_sslmode")
}
