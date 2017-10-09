package main

import (
	"github.com/juju/loggo"
	"net/http"
	"time"
)

func LogHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := loggo.GetLogger("web")
		start := time.Now()
		inner.ServeHTTP(w, r)
		logger.Tracef("%6s - %s %s %s", r.Method, r.RequestURI, name, time.Since(start))
	})
}

func OptionsHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			// TODO: Make the origin a config driven value. Currently it is effectively *
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods",
				"POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		if r.Method == "OPTIONS" {
			return
		}

		inner.ServeHTTP(w, r)
	})
}
