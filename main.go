package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stationapi/station-tag-aggregator/db"
	"github.com/stationapi/station-tag-aggregator/routes/category"
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

	r.Get("/api/tag-aggregator/tag/get", func(w http.ResponseWriter, r *http.Request) {
		err := tag.Get(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})

	r.Get("/api/tag-aggregator/tag/list", func(w http.ResponseWriter, r *http.Request) {
		err := tag.List(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})

	r.Post("/api/tag-aggregator/tag/new", func(w http.ResponseWriter, r *http.Request) {
		ok := checkAuth(r.Header.Get("Authorization"), w)

		if !ok {
			return
		}

		err := tag.Create(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})

	r.Get("/api/tag-aggregator/category/get", func(w http.ResponseWriter, r *http.Request) {
		err := category.Get(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})

	r.Get("/api/tag-aggregator/category/list", func(w http.ResponseWriter, r *http.Request) {
		err := category.List(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})

	r.Post("/api/tag-aggregator/category/new", func(w http.ResponseWriter, r *http.Request) {
		ok := checkAuth(r.Header.Get("Authorization"), w)

		if !ok {
			return
		}

		err := category.Create(w, r, db)

		if err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":3000", r)
}

func checkAuth(token string, w http.ResponseWriter) bool {
	secret := os.Getenv("ADMIN_TOKEN")

	if subtle.ConstantTimeCompare(
		[]byte(strings.Split(token, " ")[1]),
		[]byte(strings.TrimSpace(secret)),
	) != 1 {
		http.Error(w, "the provided token was incorrect", http.StatusForbidden)

		return false
	}

	return true
}
