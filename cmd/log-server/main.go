package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting log server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received") // Каждый вызов = запись в консоль
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
