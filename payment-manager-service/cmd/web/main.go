package main

import (
	"fmt"
	"payment-manager-service/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, logger)
	app := config.NewGin(viperConfig)
	validate := config.NewValidate(viperConfig)

	cfg := &config.ConfigBootstrap{
		ViperConfig: viperConfig,
		Logger:      logger,
		DB:          db,
		App:         app,
		Validate:    validate,
	}
	config.Bootstrap(cfg)

	appPort := viperConfig.GetInt("APP_PORT")
	if err := app.Run(fmt.Sprintf(":%d", appPort)); err != nil {
		logger.Fatalf("[ERROR] Failed to start Gin server : " + err.Error())
	}
}
