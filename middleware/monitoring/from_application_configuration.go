package monitoring

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
	"github.com/mtojek/greenwall/middleware/application"
)

// FromApplicationConfiguration method extracts a monitoring configuration from application configuration.
func FromApplicationConfiguration(applicationConfiguration *application.Configuration) *Configuration {
	fileContent, err := ioutil.ReadFile(applicationConfiguration.Config)
	if err != nil {
		log.Fatal(err)
	}

	var monitoringConfiguration Configuration
	err = yaml.Unmarshal(fileContent, &monitoringConfiguration)
	if err != nil {
		log.Fatal(err)
	}

	err = validate(&monitoringConfiguration)
	if err != nil {
		log.Fatal(err)
	}

	addMissingNodeType(&monitoringConfiguration) // legacy
	return &monitoringConfiguration
}

func validate(monitoringConfiguration *Configuration) error {
	if monitoringConfiguration.Groups == nil || len(monitoringConfiguration.Groups) == 0 {
		return errors.New("no groups of monitored nodes were specified")
	}
	return nil
}

func addMissingNodeType(monitoringConfiguration *Configuration) {
	for i := range monitoringConfiguration.Groups {
		for j := range monitoringConfiguration.Groups[i].Nodes {
			if len(monitoringConfiguration.Groups[i].Nodes[j].Type) == 0 {
				monitoringConfiguration.Groups[i].Nodes[j].Type = "http_check"
			}
		}
	}
}
