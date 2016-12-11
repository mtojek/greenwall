package main

import (
	"github.com/mtojek/greenwall/middleware/httpserver"
	"github.com/mtojek/greenwall/middleware/configuration"
)

func main() {
	applicationConfiguration := configuration.FromCommandLineArgs()
	httpServer := httpserver.NewHTTPServer(applicationConfiguration)
	httpServer.ListenAndServe()
}
