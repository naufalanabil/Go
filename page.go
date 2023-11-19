package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Struct untuk menyimpan data yang akan disusun ke dalam template HTML.
type PageVariables struct {
	Title   string
	Message string
}

func main() {
	// Menetapkan handler untuk rute "/".
	http.HandleFunc("/", HomePage)

	// Menetapkan handler untuk rute "/about".
	http.HandleFunc("/about", AboutPage)

	// Memulai server web pada port 8080.
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// Handler untuk rute "/".
func HomePage(w http.ResponseWriter, r *http.Request) {
	// Membuat instance dari struct PageVariables dan mengisi datanya.
	pageVariables := PageVariables{
		Title:   "Home Page",
		Message: "Welcome to the Home Page!",
	}

	// Membaca template HTML dari file.
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>{{.Title}}</title>
			</head>
			<body>
				<h1>{{.Message}}</h1>
			</body>
		</html>`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyusun template dengan data dari struct PageVariables.
	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler untuk rute "/about".
func AboutPage(w http.ResponseWriter, r *http.Request) {
	// Membuat instance dari struct PageVariables dan mengisi datanya.
	pageVariables := PageVariables{
		Title:   "About Page",
		Message: "This is the About Page!",
	}

	// Membaca template HTML dari file.
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>{{.Title}}</title>
			</head>
			<body>
				<h1>{{.Message}}</h1>
			</body>
		</html>`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyusun template dengan data dari struct PageVariables.
	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
