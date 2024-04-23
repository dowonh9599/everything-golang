package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Server")
}

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "error: 404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello World!")
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error: Parse form failed", err)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name: %s\n", name)
	fmt.Fprintf(w, "address: %s\n", address)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleGreeting)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
