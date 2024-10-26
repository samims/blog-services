package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"auth-service/controllers"
	"auth-service/db"
	"auth-service/repositories"
	"auth-service/router"
	"auth-service/services"
)

func main() {
	// DB connection
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	//repo and service intialization
	repo := repositories.NewRepository(dbConn)
	svcs := services.NewServices(repo)
	ctrl := controllers.NewController(svcs)

	r := router.InitUserRouter(ctrl)

	// create and start the HTTP server
	srv := createServer(":8080", r)

	go startServer(srv)

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

	log.Println("shuttingdown server ...")

	//  create  a new context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	//  Attempt to gracefully shutdown the server
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server forced to shutdown %v", err)
	}

	log.Println("server exited gracefully...")
}
