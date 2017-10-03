package main

import (
	"net/http"

	"github.com/juju/loggo"
	"strconv"
)

type Server struct {
}

func (s Server) ListenAndServe() error {
	cfg, err := GetConfig()
	if err != nil {
		panic(err)
	}
	LoadTemplates(cfg.Web.TemplateBaseDir)
	router := NewRouter()
	logger := loggo.GetLogger("web")

	address := ":" + strconv.Itoa(cfg.Web.Port)

	logger.Infof("Server listening on address %s", address)
	if err := http.ListenAndServe(address, router); err != nil {
		logger.Criticalf("%s", err)
	}
	return nil
}
