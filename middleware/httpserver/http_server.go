package httpserver

import (
	"net/http"

	"log"

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
	err := http.ListenAndServe(httpServer.applicationConfiguration.HostPort, staticHandler)
	if err != nil {
		log.Fatal(err)
	}
}
