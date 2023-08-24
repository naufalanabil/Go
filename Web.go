package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is the main page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
