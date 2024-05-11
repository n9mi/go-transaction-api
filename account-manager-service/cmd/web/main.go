package main

import (
	"account-manager-service/internal/config"
	"fmt"
)

func main() {
	viperCfg := config.NewViper()
	logger := config.NewLogger(viperCfg)
	db := config.NewDatabase(viperCfg, logger)
	app := config.NewGin(viperCfg)
	validate := config.NewValidate(viperCfg)
	redisClient := config.NewRedisClient(viperCfg)

	cfg := &config.ConfigBootstrap{
		ViperConfig: viperCfg,
		Log:         logger,
		DB:          db,
		App:         app,
		Validate:    validate,
		RedisClient: redisClient,
	}
	config.Bootstrap(cfg)

	appPort := viperCfg.GetInt("APP_PORT")
	if err := app.Run(fmt.Sprintf(":%d", appPort)); err != nil {
		logger.Fatalf("[ERROR] Failed to start Gin server : " + err.Error())
	}
}
