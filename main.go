package main

import (
	"github.com/davswo/BlogServices/blog"

	"github.com/davswo/BlogServices/config"
	"github.com/davswo/BlogServices/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vrischmann/envconfig"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting BlogServices...")

	var cfg config.Service
	if err := envconfig.Init(&cfg); err != nil {
		log.Panicf("Error loading main configuration %v\n", err.Error())
	}
	log.Print(cfg)

	repo, err := repository.Create(cfg.DbType)
	if err != nil {
		log.Fatal("Unable to initiate repository", err)
	}

	blogService := blog.NewBlogService(repo)

	if err := startService(cfg.Port, blogService); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

func startService(port string, blogService blog.BlogService) error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blogs", blogService.GetBlogPosts).
		Methods(http.MethodGet)

	router.HandleFunc("/user/blog", blogService.InsertBlogPost).Methods(http.MethodPost)

	log.Printf("Starting server on port %s ", port)

	c := cors.AllowAll()
	return http.ListenAndServe(":"+port, c.Handler(router))
}
