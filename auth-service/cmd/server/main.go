package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	// DB connection
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatalf("failed to close the database connection: %v", err)
		}
	}(dbConn)

	// loading env
	env := viper.New()
	appConf := config.NewAppConfig(env)
	conf := config.NewConfiguration(appConf)
	conf.Load(env)

	port := os.Getenv(constants.AppPort)

	// repo and service initialization
	repo := repositories.NewRepository(dbConn)
	svc := services.NewServices(repo, conf)
	ctrl := controllers.NewController(svc)

	r := router.InitUserRouter(ctrl)

	// create and start the HTTP server
	srv := createServer(fmt.Sprintf(":%s", port), r)

	// start the project asynchronously
	go startServer(srv)

	// shutdown
	shutdownServer(srv)

}

// createServer creates a httpServer
func createServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

func startServer(srv *http.Server) {
	log.Println("starting the server on ", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start the server: %v", err)
	}
}

// shutdownServer  shuts down the server
func shutdownServer(srv *http.Server) {
	// create  a new channel  to signal shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	//  wait for the signal
	<-stop

	log.Println("shutting down server ...")

	//  create  a new context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	//  Attempt to gracefully shut down the server
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server forced to shutdown %v", err)
	}

	log.Println("server exited gracefully...")
}
