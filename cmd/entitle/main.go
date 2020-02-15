package main

import (
	"log"
	"net/http"

	"github.com/abihf/entitle/webhook"
)

func main() {
	http.HandleFunc("/webhook", webhook.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
