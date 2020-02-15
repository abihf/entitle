package main

import (
	"log"
	"net/http"

	"github.com/abihf/entitle"
)

func main() {
	http.HandleFunc("/webhook", entitle.HandleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
