package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menetapkan handler untuk route "/".
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Memulai server web pada port 8080.
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
