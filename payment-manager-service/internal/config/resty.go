package config

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func NewRestyClient(viperCfg *viper.Viper) *resty.Client {
	return resty.New()
}
