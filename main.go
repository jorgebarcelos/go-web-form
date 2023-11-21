package main

import (
	"fmt"
	"log"
	"net/http"
	"web-form/handlers"
)


func main() {
	http.HandleFunc("/", handlers.SubscriptionHandlers)
	fmt.Println("Servindo na porta: localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}