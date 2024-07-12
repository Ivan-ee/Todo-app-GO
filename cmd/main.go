package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %s", err.Error())
	}
}
