package main

import (
	"fmt"
	"html"
	"net/http"
)

func (app *app) getHome(w http.ResponseWriter, r *http.Request) {

	// The pattern "/" matches all paths not matched by other registered routes.
	// We can use this fact to determine if the request is for an unknown route.
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: handler for %s not found", html.EscapeString(r.URL.Path))
		return
	}

	app.render(w, r, "home.page.html", nil)

}

func (app *app) getTransactions(w http.ResponseWriter, r *http.Request) {

	transactions, err := app.transactions.All()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	app.render(w, r, "transactions.page.html", pageData{"Transactions": transactions})

}

func (app *app) getAbout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the about page."))
}
