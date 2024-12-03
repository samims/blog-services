package main

import (
	"blog-service/config"
	"blog-service/controllers"
	"blog-service/db"
	"blog-service/logger"
	"blog-service/repositories"
	"blog-service/router"
	"blog-service/services"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Initialize logger
	appLogger := logger.NewAppLogger(logrus.DebugLevel)
	viperEnv, err := setupEnv()
	if err != nil {
		appLogger.Fatal(err)
		return
	}

	// Application config
	appConfig := config.NewAppConfig(viperEnv)

	// PostGreSQL config
	postgresConfig := config.NewPostgresConfig(viperEnv)
	postgresConnector := db.NewPostgresConnector(postgresConfig, appLogger)

	ctx := context.Background()
	dbConn, err := postgresConnector.Connect(ctx)
	if err != nil {
		appLogger.WithContext(ctx).Fatal(err)
	}

	// Initialize repository, service, and controller
	blogRepo := repositories.NewBlogRepository(dbConn, appLogger)
	blogService := services.NewBlogService(blogRepo, appLogger)
	blogController := controllers.NewBlogController(blogService, appLogger)

	// Initialize router
	r := router.Init(blogController)
	appLogger.Infof("Starting server on port :%s", appConfig.GetPort())

	// Start server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", appConfig.GetPort()), r); err != nil {
		appLogger.Fatal(err)
	}
}

func getCurrDir() string {
	// show me the current directory name
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir

}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func setupEnv() (*viper.Viper, error) {
	viperEnv := viper.New()

	// .env file location
	envPath := filepath.Join(getCurrDir(), ".env")
	if !isFileExist(envPath) {
		viperEnv.AutomaticEnv()
		return viperEnv, nil
	}
	// load from .env file
	viperEnv.SetConfigFile(envPath)
	err := viperEnv.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viperEnv, nil

}
