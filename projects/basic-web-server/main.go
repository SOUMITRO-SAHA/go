package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		fmt.Fprintf(w, "Please enter a name and email")
		return
	}

	fmt.Fprintf(w, "Form Submitted Successfully!")

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Email: %s\n", email)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	// Routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Staring server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
