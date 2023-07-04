package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	router := http.NewServeMux()
	router.HandleFunc("/", handler)
	router.HandleFunc("/new", newHandler)

	log.Println("Server in running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Monolithic app: original handler")
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New functionality: new handler")
}

