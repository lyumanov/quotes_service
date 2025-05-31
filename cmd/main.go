package main

import (
	"fmt"
	"net/http"
	"quotes_service/internal/handler"
	"quotes_service/internal/repo"
	"quotes_service/internal/routes"
)

func main() {

	repository := repo.NewQuoteRepo()
	handlers := handler.NewQuoteHandler(repository)
	router := routes.RegisterRoute(handlers)

	fmt.Println("Сервер запущен на 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Не удалось запустить сервер")
	}
}
