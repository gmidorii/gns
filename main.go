package main

import (
	"log"
	"net/http"
)

func run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})
	return http.ListenAndServe(":53", nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
