package main

import (
	"log"
	"net/http"

	service "github.com/Peyton232/dual-go-reciever/services"
)

func main() {
	log.Println("server started")
	http.HandleFunc("/webhookIridium", service.HandleWebhook)
	http.HandleFunc("/webhookLora", service.HandleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
