package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("Server is running!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
