package main

import (
	"github.com/mtojek/greenwall/middleware/application"
	"github.com/mtojek/greenwall/middleware/healthcheck"
	"github.com/mtojek/greenwall/middleware/httpserver"
	"github.com/mtojek/greenwall/middleware/monitoring"
)

func main() {
	applicationConfiguration := application.ReadConfiguration()
	monitoringConfiguration := monitoring.FromApplicationConfiguration(applicationConfiguration)
	healthcheck := healthcheck.NewHealthcheck(applicationConfiguration, monitoringConfiguration)
	healthcheck.Start()
	httpServer := httpserver.NewHTTPServer(applicationConfiguration, monitoringConfiguration, healthcheck)
	httpServer.ListenAndServe()
}
