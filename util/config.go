/*
utils contains functions for misl tasks
*/
package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port      string `mapstructure:"SERVER_PORT"`
	DBSource  string `mapstructure:"DB_SOURCE"`
	DBDriver  string `mapstructure:"DB_DRIVER"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

// LoadConfig takes relative route to app.env from a file and returns env variables stored in it
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
