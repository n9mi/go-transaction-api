package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viperConfig *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := viperConfig.GetString("SERVICE_DB_HOST")
	user := viperConfig.GetString("SERVICE_DB_USER")
	password := viperConfig.GetString("SERVICE_DB_PASSWORD")
	name := viperConfig.GetString("SERVICE_DB_NAME")
	port := viperConfig.GetInt("SERVICE_DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		password,
		name,
		port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             5 * time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		log.Fatalf("failed to connect into database: %+v", err)
	}

	_, err = db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection : %+v", err)
	}

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (w *logrusWriter) Printf(message string, args ...interface{}) {
	w.Logger.Tracef(message, args...)
}
