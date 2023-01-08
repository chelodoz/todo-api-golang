package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	MongoHost         string `mapstructure:"MONGO_HOST"`
	MongoPort         string `mapstructure:"MONGO_PORT"`
	MongoDatabase     string `mapstructure:"MONGO_DATABASE"`
	MongoCollection   string `mapstructure:"MONGO_COLLECTION"`
	MongoUsername     string `mapstructure:"MONGO_USERNAME"`
	MongoPassword     string `mapstructure:"MONGO_PASSWORD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
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
