package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/husseinamine/florasrv/views"
)

func main() {
	logger := log.New(os.Stdout, "[FLORA] ", log.LstdFlags)

	smux := http.NewServeMux()
	users := views.NewUsers(logger)

	smux.Handle("/users/", users)

	server := &http.Server{
		Addr:     ":8080",
		ErrorLog: logger,
		Handler:  smux,

		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)
	signal.Notify(shutdown, os.Kill)

	<-shutdown
	logger.Println("Graceful Shutdown!")

	sc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(sc)
	cancel()
}
