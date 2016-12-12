package main

import (
	"github.com/mtojek/greenwall/middleware/application"
	"github.com/mtojek/greenwall/middleware/httpserver"
	"github.com/mtojek/greenwall/middleware/monitoring"
)

func main() {
	applicationConfiguration := application.FromCommandLineArgs()
	monitoringConfiguration := monitoring.FromApplicationConfiguration(applicationConfiguration)
	httpServer := httpserver.NewHTTPServer(applicationConfiguration, monitoringConfiguration)
	httpServer.ListenAndServe()
}
