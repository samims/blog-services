package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"auth-service/config"
	"auth-service/constants"
	"auth-service/controllers"
	"auth-service/db"
	"auth-service/repositories"
	"auth-service/router"
	"auth-service/services"

	"github.com/spf13/viper"
)

func main() {
	dbConn := mustConnectDB()
	defer mustCloseDB(dbConn)

	l := logrus.New()
	env := loadEnv()
	conf := config.NewConfiguration(config.NewAppConfig(env))

	repo := mustInitRepo(dbConn)
	svc := services.NewServices(repo, conf)
	ctrl := controllers.NewController(svc, l)

	r := router.InitUserRouter(ctrl)
	srv := createServer(fmt.Sprintf("0.0.0.0:%s", os.Getenv(constants.AppPort)), r)

	go startServer(srv)
	shutdownServerGracefully(srv)
}

// mustConnectDB establishes a database connection and panics if it fails.
func mustConnectDB() *sql.DB {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	return dbConn
}

// mustCloseDB closes the database connection and panics if it fails.
func mustCloseDB(dbConn *sql.DB) {
	if err := dbConn.Close(); err != nil {
		log.Fatalf("failed to close the database connection: %v", err)
	}
}

// loadEnv loads environment variables and panics if it fails.
func loadEnv() *viper.Viper {
	env := viper.New()
	env.AutomaticEnv()
	return env
}

// mustInitRepo initializes the repository and panics if it fails.
func mustInitRepo(dbConn *sql.DB) repositories.Repository {
	repo, err := repositories.NewRepository(dbConn)
	if err != nil {
		log.Fatalf("failed to initialize the repository: %v", err)
	}
	return repo
}

// createServer creates a configured HTTP server.
func createServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

// startServer starts the HTTP server and logs fatal errors.
func startServer(srv *http.Server) {
	log.Println("starting server on..", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start the server: %v", err)
	}
}

// shutdownServerGracefully handles graceful server shutdown on interrupt signal.
func shutdownServerGracefully(srv *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exited gracefully")
}
