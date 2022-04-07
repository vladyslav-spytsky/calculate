package main

import (
	"calculate/conf"
	"calculate/src/domain/infrastructure/delivery"
	"calculate/src/registry"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func main() {
	// Init configs
	conf.InitConfig()

	// New router
	router := httprouter.New()

	r := registry.NewRegistry()

	// Init router
	router = delivery.NewRouter(router, r.NewAPIController())

	err := http.ListenAndServe(viper.GetString("http.port"), router)
	if err != nil {
		panic(err)
	}
}
