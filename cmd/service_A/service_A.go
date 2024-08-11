package main

import (
	"go-microservices-template/pkg/service_A/controllers"
	"log"
	"net/http"
)

func main() {
	controller, err := controllers.NewController()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", controller.Ping)

	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
