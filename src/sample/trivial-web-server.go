package main

import (
	"fmt"
	"net/http"
	"math/rand"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am a GO application running inside Docker.")

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
    min := 0
    max := 100
	fmt.Fprintf(w, rand.Intn(max - min) + min)
}

func main() {
	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/random", randomHandler)
	http.ListenAndServe(":8080", nil)
}
