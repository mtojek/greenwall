package checks

import (
	"reflect"

	"log"

	"github.com/mtojek/greenwall/middleware/monitoring"
)

const (
	// MessageNotCheckedYet means that no checks have been run yet.
	MessageNotCheckedYet = "Not checked yet"
	// MessagePatternNotFound means that the searched pattern has not been found.
	MessagePatternNotFound = "Pattern not found"
	// MessageOK means a successful check.
	MessageOK = "OK"

	// StatusDanger set the background of node tile to red.
	StatusDanger = "danger"
	// StatusSuccess set the background of node tile to green.
	StatusSuccess = "success"
	// StatusWarning set the background of node tile to yellow.
	StatusWarning = "warning"
)

// Check describes the basic interface of every health check.
type Check interface {
	Initialize(monitoringConfiguration *monitoring.Configuration, nodeConfig *monitoring.Node)
	Run() CheckResult
}

var (
	checkTypeRegistry = map[string]reflect.Type{}
)

// CheckResult represents a result of an executed health check.
type CheckResult struct {
	Status  string
	Message string
}

func registerType(name string, aType reflect.Type) {
	checkTypeRegistry[name] = aType
}

// MakeInstance method creates a new instance of check based on a given name.
func MakeInstance(name string) Check {
	aType, ok := checkTypeRegistry[name]
	if !ok {
		log.Fatalf("Unknown check type: %s", name)
	}
	return reflect.New(aType).Interface().(Check)
}
