package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dbsimmons64/go-beans/internal"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	templates    TemplateCache
	transactions *internal.TransactionModel
}

func main() {

	// Connect to the sqlite database
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	// Cache any templates
	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		templates: templateCache,
		transactions: &internal.TransactionModel{
			DB: db,
		},
	}

	srv := http.Server{
		Addr:    ":8081",
		Handler: app.routes(),
	}

	log.Println("Listening on port", srv.Addr)
	srv.ListenAndServe()
}
