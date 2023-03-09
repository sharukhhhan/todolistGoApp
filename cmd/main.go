package main

import (
	todo "To-do-list"
	"To-do-list/pkg/handler"
	"To-do-list/pkg/repository"
	"To-do-list/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred during running the http server : %s", err.Error())
	}

}
