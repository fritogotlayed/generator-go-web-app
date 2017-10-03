package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.Handler
		handler = LogHandler(handler, route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	// TODO: If the resources ever get embedded this needs to change.
	handler := LogHandler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))), "static")
	router.PathPrefix("/static/").Handler(handler)

	return router
}
