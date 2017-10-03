package main

import (
	"github.com/juju/loggo"
)

func main() {
	// Levels: TRACE, DEBUG, INFO, WARNING, ERROR, CRITICAL
	loggo.ConfigureLoggers(`<root>=TRACE; main=TRACE`)
	logger := loggo.GetLogger("main")
	logger.Infof("Application starting...")

	server := Server{}
	server.ListenAndServe()

	logger.Infof("Application Exiting.")
}
