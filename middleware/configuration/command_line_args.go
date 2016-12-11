package configuration

import "flag"

func FromCommandLineArgs() *ApplicationConfiguration {
	hostPort := flag.String("hostPort", ":9001", "Host:port of the greenwall HTTP server")
	staticDir := flag.String("staticDir", "frontend", "Path to frontend static resources")
	flag.Parse()

	return &ApplicationConfiguration{
		HostPort:  *hostPort,
		StaticDir: *staticDir,
	}
}
