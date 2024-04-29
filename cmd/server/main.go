package main

import (
	"context"
	"joi-energy-golang/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	Run()
}

// Run starts the HTTP server
func Run() {
	server := router.NewServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		println()
		log.Println("Shutting down server...")

		err := gracefulShutdown(server, 25*time.Second)

		if err != nil {
			log.Printf("Server stopped: %s", err.Error())
		}

		os.Exit(0)
	}()

	log.Printf("Listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func gracefulShutdown(server *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return server.Shutdown(ctx)
}
