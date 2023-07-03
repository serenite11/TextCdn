package main

import (
	"Test-Task/TextCDN/internal/app/rest"
	"Test-Task/TextCDN/internal/app/rest/handlers"
	"Test-Task/TextCDN/internal/database"
	"Test-Task/TextCDN/internal/repository"
	"Test-Task/TextCDN/internal/services"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	collection := database.InitMongoDb()
	repos := repository.NewRepository(collection)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)
	server := new(rest.Server)
	go func() {
		if err := server.Run(":8080", handler.InitRoutes()); err != nil {
			log.Fatalf(err.Error())
			return
		}
	}()
	log.Println("App is started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.ShuttingDown(context.Background()); err != nil {
		log.Fatalf(err.Error())
		return
	}
}
