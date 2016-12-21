package checks

import (
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/mtojek/greenwall/middleware/monitoring"
)

const (
	sampleCheckName    = "sample_check"
	messageNotGreenDay = "Not a green day!"
	greenDayParameter  = "greenDay"
)

// SampleCheck is a sample health check, comparing the current day with a "green" day.
type SampleCheck struct {
	monitoringConfiguration *monitoring.Configuration
	nodeConfig              *monitoring.Node
}

// Initialize method initializes the check instance.
func (s *SampleCheck) Initialize(monitoringConfiguration *monitoring.Configuration, nodeConfig *monitoring.Node) {
	s.monitoringConfiguration = monitoringConfiguration
	s.nodeConfig = nodeConfig
}

// Run method executes the check. This is invoked periodically.
func (s *SampleCheck) Run() CheckResult {
	greenDay, err := strconv.Atoi(s.nodeConfig.Parameters[greenDayParameter])
	if err != nil {
		log.Fatal(err) // This should never happen.
	}

	if time.Now().Day() == greenDay {
		return CheckResult{
			Message: MessageOK,
			Status:  StatusSuccess,
		}
	}
	return CheckResult{
		Message: messageNotGreenDay,
		Status:  StatusDanger,
	}
}

func init() {
	registerType(sampleCheckName, reflect.TypeOf(SampleCheck{}))
}
