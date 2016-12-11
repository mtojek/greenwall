package httpserver

import (
	"net/http"

	"github.com/mtojek/greenwall/middleware/configuration"
)

type HTTPServer struct {
	applicationConfiguration *configuration.ApplicationConfiguration
}

func NewHTTPServer(applicationConfiguration *configuration.ApplicationConfiguration) *HTTPServer {
	return &HTTPServer{
		applicationConfiguration: applicationConfiguration,
	}
}

func (httpServer *HTTPServer) ListenAndServe() {
	staticHandler := http.FileServer(http.Dir(httpServer.applicationConfiguration.StaticDir))
	http.ListenAndServe(httpServer.applicationConfiguration.HostPort, staticHandler)
}
