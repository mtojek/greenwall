package httpserver

import (
	"net/http"

	"log"

	"github.com/mtojek/greenwall/middleware/configuration"
)

// HTTPServer is responsible for serving HTTP requests for frontend resources.
type HTTPServer struct {
	applicationConfiguration *configuration.ApplicationConfiguration
}

// NewHTTPServer method creates new instance of HTTPServer.
func NewHTTPServer(applicationConfiguration *configuration.ApplicationConfiguration) *HTTPServer {
	return &HTTPServer{
		applicationConfiguration: applicationConfiguration,
	}
}

// ListenAndServe method listens and serves requests sent to HTTP handlers.
func (httpServer *HTTPServer) ListenAndServe() {
	staticHandler := http.FileServer(http.Dir(httpServer.applicationConfiguration.StaticDir))
	err := http.ListenAndServe(httpServer.applicationConfiguration.HostPort, staticHandler)
	if err != nil {
		log.Fatal(err)
	}
}
