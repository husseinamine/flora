package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/husseinamine/florasrv/apps"
	"github.com/husseinamine/florasrv/routes"
)

func main() {
	logger := log.New(os.Stdout, "[FLORA] ", log.LstdFlags)

	smux := mux.NewRouter()
	users := apps.NewUsers(logger)

	routes.NewUsers(smux, users).Initialize()

	// SERVER CONFIGURATION
	server := &http.Server{
		Addr:     ":8080",
		ErrorLog: logger,
		Handler:  smux,

		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// START LISTENING PROCESS
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()

	// GRACEFUL SHUTDOWN PROCEDURES
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)
	signal.Notify(shutdown, os.Kill)

	<-shutdown
	logger.Println("Graceful Shutdown!")

	sc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(sc)
}
