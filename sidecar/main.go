package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from main app")
}

func sidecardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Doing task in the sidecar")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", mux))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(sidecardHandler)))
	}()

	select {}
}
