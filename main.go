package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", handleHook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
