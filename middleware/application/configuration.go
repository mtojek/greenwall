package application

import (
	"flag"
	"log"
	"os"
	"path"

	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const indexFile = "index.html"

var (
	config    = flag.String("config", "config.yaml", "A config file defining monitored nodes (env: CONFIG)")
	hostPort  = flag.String("hostPort", ":9001", "Host:port of the greenwall HTTP server (env: HOST, PORT)")
	staticDir = flag.String("staticDir", "frontend", "Path to frontend static resources (env: STATIC_DIR)")
)

// Configuration represents a required GreenWall configuration.
type Configuration struct {
	Config    string
	HostPort  string
	StaticDir string
}

type environmentConfiguration struct {
	Config    string
	Host      string
	Port      int
	StaticDir string `envconfig:"STATIC_DIR"`
}

// ReadConfiguration method reads configuration from chosen source.
func ReadConfiguration() *Configuration {
	flag.Parse()

	ec := new(environmentConfiguration)
	err := envconfig.Process("", ec)
	if err != nil {
		log.Fatal(err)
	}

	c := new(Configuration)
	if len(ec.Config) > 0 {
		c.Config = ec.Config
	} else {
		c.Config = *config
	}

	if len(ec.Host) > 0 || ec.Port > 0 {
		c.HostPort = fmt.Sprintf("%s:%d", ec.Host, ec.Port)
	} else {
		c.HostPort = *hostPort
	}

	if len(ec.StaticDir) > 0 {
		c.StaticDir = ec.StaticDir
	} else {
		c.StaticDir = *staticDir
	}

	err = validate(c)
	if err != nil {
		log.Fatalf("Error occurred while validating configuration: %v", err)
	}
	return c
}

func validate(applicationConfiguration *Configuration) error {
	indexFile := path.Join(applicationConfiguration.StaticDir, indexFile)
	_, err := os.Stat(indexFile)
	if err != nil {
		return err
	}
	_, err = os.Stat(applicationConfiguration.Config)
	return err
}
