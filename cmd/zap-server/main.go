package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Сброс буфера при завершении

	fmt.Println("Starting zap server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received") // Буферизованный вывод
		w.Write([]byte("OK"))
	})

	logger.Fatal("Server crashed",
		zap.Error(http.ListenAndServe(":8081", nil)),
	)
}
