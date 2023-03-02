package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stationapi/station-tag-aggregator/db"
	"github.com/stationapi/station-tag-aggregator/routes/tag"
)

func main() {
	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)	
	r.Use(middleware.RealIP)

	r.Get("/api/tag/get", func(w http.ResponseWriter, r *http.Request) {
		err := tag.Get(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})

	r.Post("/api/tag/new", func(w http.ResponseWriter, r *http.Request) {
		err := tag.Create(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})
}
