package main

import (
	"github.com/gorilla/mux"
)

// routes returns a router with all paths
func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(app.metricsAndHealthMiddleware)

	r.HandleFunc("/slow", app.slow).
		Methods("GET")
	r.HandleFunc("/configure/liveness", app.switchLiveness).
		Methods("GET")
	r.HandleFunc("/configure/readiness", app.switchReadiness).
		Methods("GET")
	r.PathPrefix("/").HandlerFunc(app.home).
		Methods("GET")

	return r
}
