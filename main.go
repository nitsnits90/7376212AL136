package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
	http.HandleFunc("/harini", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, harini!")
	})
	http.HandleFunc("/212AL115", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 212AL115!")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "red!")
	})
	http.HandleFunc("/school", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "st.joseph")
		
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
