package main

import (
	delivery2 "calculate/src/infrastructure/delivery"
	"calculate/src/registry"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// New router
	router := httprouter.New()

	r := registry.NewRegistry()

	// Init router
	router = delivery2.NewRouter(router, r.NewAPIController())

	err := http.ListenAndServe(":8989", router)
	if err != nil {
		panic(err)
	}
}
