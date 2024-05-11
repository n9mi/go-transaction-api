package config

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigType("env")
	config.SetConfigFile(".env")
	config.AddConfigPath("./")
	config.AddConfigPath("./../")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	return config
}
