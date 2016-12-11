package main

import (
	"github.com/mtojek/greenwall/middleware/configuration"
	"github.com/mtojek/greenwall/middleware/httpserver"
)

func main() {
	applicationConfiguration := configuration.FromCommandLineArgs()
	httpServer := httpserver.NewHTTPServer(applicationConfiguration)
	httpServer.ListenAndServe()
}
