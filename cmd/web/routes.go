package main

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.getHome)
	mux.HandleFunc("GET /transactions", app.getTransactions)
	mux.HandleFunc("GET /about", app.getAbout)

	return mux
}
