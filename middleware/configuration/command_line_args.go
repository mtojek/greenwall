package configuration

import (
	"flag"
	"log"
	"os"
	"path"
)

const indexFile = "index.html"

// FromCommandLineArgs method reads configuration from command line arguments.
func FromCommandLineArgs() *ApplicationConfiguration {
	hostPort := flag.String("hostPort", ":9001", "Host:port of the greenwall HTTP server")
	staticDir := flag.String("staticDir", "frontend", "Path to frontend static resources")
	flag.Parse()

	applicationConfiguration := &ApplicationConfiguration{
		HostPort:  *hostPort,
		StaticDir: *staticDir,
	}

	err := validate(applicationConfiguration)
	if err != nil {
		log.Fatalf("Error occurred while validating configuration: %v", err)
	}
	return applicationConfiguration
}

func validate(applicationConfiguration *ApplicationConfiguration) error {
	indexFile := path.Join(applicationConfiguration.StaticDir, indexFile)
	_, err := os.Stat(indexFile)
	return err
}
