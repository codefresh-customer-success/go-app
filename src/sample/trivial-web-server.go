package main

import (
	"fmt"
	"strconv"
	"net/http"
	"math/rand"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am a GO application running inside Docker.")

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
    min := 1
    max := 30
	fmt.Fprintf( w, strconv.Itoa( rand.Intn(max - min + 1) + min ) ) 
}

func main() {
	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/random", randomHandler)
	http.ListenAndServe(":8080", nil)
}
