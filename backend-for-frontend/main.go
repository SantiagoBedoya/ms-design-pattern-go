package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/bff", bffHandler)
	http.ListenAndServe(":8080", nil)

}

func bffHandler(w http.ResponseWriter, r *http.Request) {
	clientType := r.Header.Get("Client-Type")

	var backendResponse string
	switch clientType {
	case "web":
		backendResponse = callWebBackend()
	case "mobile":
		backendResponse = callMobileBackend()
	default:
		http.Error(w, "Client Type no compatible", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, backendResponse)
}

func callWebBackend() string {
	return "Backend Web response"
}

func callMobileBackend() string {
	return "Backend Mobile response"
}
