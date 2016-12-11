package httpserver

import "net/http"

// ServerMux is responsible for routing requests to appropriate handlers.
type ServerMux struct {
	indexHandler  *IndexHandler
	staticHandler http.Handler
}

// NewServerMux method creates a new instance of ServerMux.
func NewServerMux(indexHandler *IndexHandler, staticHandler http.Handler) *ServerMux {
	return &ServerMux{
		indexHandler:  indexHandler,
		staticHandler: staticHandler,
	}
}

// ServeHTTP method performs routing of requests to appropriate handlers.
func (serverMux *ServerMux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" || req.URL.Path == "/index.html" {
		serverMux.indexHandler.ServeHTTP(rw, req)
		return
	}
	serverMux.staticHandler.ServeHTTP(rw, req)
}
