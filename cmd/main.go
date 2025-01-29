package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/julessinger/ark-validator/internal/application"
)

func main() {
	log.Println("Starting application...")

	service := application.New()
	service.Start()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	service.Stop()
	log.Println("Application stopped")
}
