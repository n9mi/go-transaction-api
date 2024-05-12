package main

import (
	"fmt"
	"payment-manager-service/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	app := config.NewGin(viperConfig)
	validate := config.NewValidate(viperConfig)
	resty := config.NewRestyClient(viperConfig)

	cfg := &config.ConfigBootstrap{
		ViperConfig: viperConfig,
		Log:         log,
		DB:          db,
		App:         app,
		Validate:    validate,
		RestyClient: resty,
	}
	config.Bootstrap(cfg)

	appPort := viperConfig.GetInt("APP_PORT")
	if err := app.Run(fmt.Sprintf(":%d", appPort)); err != nil {
		log.Fatalf("[ERROR] Failed to start Gin server : " + err.Error())
	}
}
