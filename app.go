package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
    healthzPath  = "/probe/liveness"
    readinessPath = "/probe/readiness"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello Filip")
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Healthy!")
}

func readyz(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Ready!")
}

func handleRequests() {
	http.HandleFunc("/", hello)
	http.HandleFunc(healthzPath, healthz)
	http.HandleFunc(readinessPath, readyz)

	port := "11000"
	
	// if port == "" {
	// 	log.Fatal("Port is not set.")
	// }

	log.Printf("Server listening on port %s...", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func main() {
	handleRequests()
}