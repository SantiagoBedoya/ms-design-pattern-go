package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	consul "github.com/hashicorp/consul/api"
)

func main() {
	config := consul.DefaultConfig()
	client, err := consul.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	
	serviceID := "my-service"
	serviceName := "my-service"
	servicePort := 8080
	serviceTags := []string{"tag1", "tag2"}

	registration := new(consul.AgentServiceRegistration)
	registration.ID = serviceID
	registration.Name = serviceName
	registration.Port = servicePort
	registration.Tags = serviceTags

	addr, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	registration.Address = addr

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}



	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		err := client.Agent().ServiceDeregister(serviceID)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from service")
	})

	log.Printf("Service is running on port %d...", servicePort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil))
}
