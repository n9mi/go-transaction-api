package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viperConfig *viper.Viper) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.Level(viperConfig.GetInt32("LOG_LEVEL")))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
