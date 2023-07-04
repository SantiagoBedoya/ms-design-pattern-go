package main


import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var backendURLs = map[string]string{
	"service1": "http://localhost:8080",
	"service2": "http://localhost:8081",
	"service3": "http://localhost:8082",
}

func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Path[1:]

	backendURL, ok := backendURLs[serviceName]
	if !ok {
		http.NotFound(w, r)
		return
	}

	targetURL, err := url.Parse(backendURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(w, r)
}

func main(){
	http.HandleFunc("/", gatewayHandler)

	fmt.Println("API Gateway running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
