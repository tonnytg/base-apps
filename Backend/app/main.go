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

	// ListenAndServe listens on the TCP network address addr and then calls
	// Serve with handler to handle requests on incoming connections.
	if os.Getenv("PORT") != "" {
		log.Println("Listening on port " + os.Getenv("PORT"))
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	}
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
