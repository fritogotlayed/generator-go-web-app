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
