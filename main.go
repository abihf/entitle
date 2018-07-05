package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// stupid hack to make
	go keepAlive()

	http.HandleFunc("/webhook", handleHook)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "PONG")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func keepAlive() {
	for {
		time.Sleep(10 * time.Second)
		http.Get("https://entitle.now.sh/ping")
	}
}
