package main

import (
	"log"

	"github.com/dkeyceo/todo-app"
	"github.com/dkeyceo/todo-app/pkg/handler"
	"github.com/dkeyceo/todo-app/pkg/repository"
	"github.com/dkeyceo/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8099", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured: %s", err.Error())
	}
}
