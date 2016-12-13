package httpserver

import (
	"log"
	"net/http"

	"github.com/mtojek/greenwall/middleware/application"
	"github.com/mtojek/greenwall/middleware/healthcheck"
	"github.com/mtojek/greenwall/middleware/monitoring"
)

// HTTPServer is responsible for serving HTTP requests for frontend resources.
type HTTPServer struct {
	applicationConfiguration *application.Configuration
	serverMux                *ServerMux
}

// NewHTTPServer method creates new instance of HTTPServer.
func NewHTTPServer(applicationConfiguration *application.Configuration,
	monitoringConfiguration *monitoring.Configuration, healthcheck *healthcheck.Healthcheck) *HTTPServer {
	indexHandler := NewIndexHandler(applicationConfiguration, monitoringConfiguration, healthcheck)
	staticHandler := http.FileServer(http.Dir(applicationConfiguration.StaticDir))
	serverMux := NewServerMux(indexHandler, staticHandler)
	return &HTTPServer{
		applicationConfiguration: applicationConfiguration,
		serverMux:                serverMux,
	}
}

// ListenAndServe method listens and serves requests sent to HTTP handlers.
func (httpServer *HTTPServer) ListenAndServe() {
	err := http.ListenAndServe(httpServer.applicationConfiguration.HostPort, httpServer.serverMux)
	if err != nil {
		log.Fatal(err)
	}
}
