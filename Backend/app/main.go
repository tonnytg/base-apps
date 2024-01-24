package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// HandleFunc registers the handler function for the given pattern
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Println("Listening on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
