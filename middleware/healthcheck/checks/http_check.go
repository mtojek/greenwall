package checks

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/mtojek/greenwall/middleware/monitoring"
)

const (
	httpCheckName            = "http_check"
	expectedPatternParameter = "expectedPattern"
)

// HTTPCheck performs a simple check against an HTTP endpoint.
type HTTPCheck struct {
	monitoringConfiguration *monitoring.Configuration
	nodeConfig              *monitoring.Node

	client          *http.Client
	searchedPattern []byte
}

// Initialize method initializes the check instance.
func (h *HTTPCheck) Initialize(monitoringConfiguration *monitoring.Configuration, nodeConfig *monitoring.Node) {
	h.monitoringConfiguration = monitoringConfiguration
	h.nodeConfig = nodeConfig

	h.client = &http.Client{Timeout: h.monitoringConfiguration.General.HTTPClientTimeout}

	h.searchedPattern = []byte(h.nodeConfig.ExpectedPattern) // Deprecation
	if len(h.searchedPattern) == 0 {
		h.searchedPattern = []byte(h.nodeConfig.Parameters[expectedPatternParameter])
	}
}

// Run method executes the check. This is invoked periodically.
func (h *HTTPCheck) Run() CheckResult {
	result := CheckResult{
		Status: StatusDanger,
	}

	response, err := h.client.Get(h.nodeConfig.Endpoint)
	if err != nil {
		log.Println(err)

		result.Message = err.Error()
		return result
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)

		result.Message = err.Error()
		return result
	}

	if response.Body != nil {
		errClosing := response.Body.Close()
		if err != nil {
			log.Println(errClosing)

			result.Message = errClosing.Error()
			return result
		}
	}

	if len(h.searchedPattern) > 0 && !bytes.Contains(body, h.searchedPattern) {
		result.Message = MessagePatternNotFound
		return result
	}

	result.Status = StatusSuccess
	result.Message = MessageOK
	return result
}

func init() {
	registerType(httpCheckName, reflect.TypeOf(HTTPCheck{}))
}
