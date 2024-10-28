package main

import (
	"blog-service/config"
	"blog-service/controllers"
	"blog-service/db"
	"blog-service/logger"
	"blog-service/repositories"
	"blog-service/services"
	"context"
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// init log
	l := logger.Log
	l.WithFields(logrus.Fields{
		"version": "0.1.1",
		"service": "blog-service",
	})

	viperEnv := viper.New()
	conf := config.NewAppConfig(viperEnv)
	// pg config
	pgConf := config.NewPostgresConfig(viperEnv)

	pgConnector := db.NewPostgresConnector(pgConf, l)
	ctx := context.Background()
	dbConn, err := pgConnector.Connect(ctx)
	if err != nil {
		l.Fatal(err)
	}

	// get the repo
	blogRepo := repositories.NewBlogRepository(dbConn, l)
	blogService := services.NewBlogService(blogRepo, l)
	blogController := controllers.NewBlogController(blogService, l)

	_ = blogController

	r := http.NewServeMux()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// run the server
	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.GetBuildEnv()), r)
	if err != nil {
		l.Fatal(err)
		return
	}
}
