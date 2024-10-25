package main

import (
	"blog-service/config"
	"blog-service/db"
	"blog-service/logger"
	"context"
	"database/sql"
	"time"

	"github.com/spf13/viper"
)

func main() {
	cfg := config.NewAppConfig(viper.New())
	l := logger.Log
	ctx := context.Background()
	connector := db.NewBaseConnector(cfg, l)

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	database, err := connector.Connect(ctx)
	if err != nil {
		l.Errorf("failed to connect %s", err.Error())
		return
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			l.Errorf("failed to close database %s", err.Error())
		}
	}(database)

}
