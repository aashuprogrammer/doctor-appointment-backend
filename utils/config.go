package utils

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbUrl               string        `mapstructure:"DB_URL"`
	DbDriver            string        `mapstructure:"DB_DRIVER"`
	Port                string        `mapstructure:"PORT"`
	SecretKey           string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("unable to decode", err)
		return
	}
	return config, err
}
