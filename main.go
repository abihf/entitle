package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// stupid hack to make
	keepAliveURL := os.Getenv("KEEP_ALIVE_URL")
	if keepAliveURL != "" {
		go keepAlive(keepAliveURL)
	}

	http.HandleFunc("/webhook", handleHook)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "PONG")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func keepAlive(url string) {
	for {
		time.Sleep(10 * time.Second)
		http.Get(url)
	}
}
