package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am a GO application running inside Docker.")

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}

func main() {
	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.ListenAndServe(":8080", nil)
}
