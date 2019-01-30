package main

import (
	"encoding/json"
	"github.com/davswo/BlogServices/blog"
	"github.com/davswo/BlogServices/config"
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

	if err := startService(cfg.Port); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

func startService(port string) error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blogs", getStaticBlogEntries).
		Methods(http.MethodGet)

	log.Printf("Starting server on port %s ", port)

	c := cors.AllowAll()
	return http.ListenAndServe(":"+port, c.Handler(router))
}

func getStaticBlogEntries(w http.ResponseWriter, r *http.Request) {
	blogEntries := blog.BlogEntries{
		{Title: "O’zapft is!",
			Text: "Bavaria ipsum dolor sit amet Zidern so, " +
				"iabaroi Mamalad da in da. Gwihss mehra barfuaßat großherzig singd a am acht’n Tag schuf Gott des Bia. " +
				"Do des back mas, sog i schnacksln. Hallelujah sog i, luja Wiesn bittschön a fescha Bua sodala auf der Oim, " +
				"da gibt’s koa Sünd is ma Wuascht kimmt i Biaschlegl: Wann griagd ma nacha wos z’dringa umma nomoi mei Weiznglasl " +
				"jo mei is des schee aba des basd scho ognudelt mechad: Kuaschwanz Habedehre hawadere midananda obandln, " +
				"des wiad a Mordsgaudi kumm geh! Ned Xaver a ganze Hoiwe des is a gmahde Wiesn nix Gwiass woass ma ned unbandig Watschnbaam" +
				" is des liab Wiesn mi. Gscheckate ned is ma Wuascht, Steckerleis. Obandln Baamwach hogg di hera am acht’n Tag schuf " +
				"Gott des Bia nia is, ognudelt pfiad de nackata Wurschtsolod. Allerweil Watschnbaam Schbozal auszutzeln nois Bussal von Greichats. " +
				"Oba da jo mei is des schee, .",
			Author: "Huaba Bua",
		},
		{Title: "Ja kurze fix!",
			Text: "Bavaria ipsum dolor sit amet fias Bradwurschtsemmal, i mechad dee Schwoanshaxn. " +
				"Nia need Stubn fensdaln a ganze, Bussal pfenningguat. " +
				"I daad wuid fei Biawambn mogsd a Bussal Goaßmaß scheans naa aasgem i.",
			Author: "Goaßl Sepp",
		},
	}
	json.NewEncoder(w).Encode(blogEntries)
}
