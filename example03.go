package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from webserver GO")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "good afternoon guy")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe("localhost:8000", nil)
}
