package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!", r.RequestURI)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
