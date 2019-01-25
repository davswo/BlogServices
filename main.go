package main

import (
	"github.com/davswo/BlogServices/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vrischmann/envconfig"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting BlogWebServices...")

	var cfg config.Service
	if err := envconfig.Init(&cfg); err != nil {
		log.Panicf("Error loading main configuration %v\n", err.Error())
	}
	log.Print(cfg)

	if err := startService(cfg.Port); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

func startService(port string) error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blogs",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Some Blog Posts will be sent here\n"))
		}).
		Methods(http.MethodGet)

	log.Printf("Starting server on port %s ", port)

	c := cors.AllowAll()
	return http.ListenAndServe(":"+port, c.Handler(router))
}
