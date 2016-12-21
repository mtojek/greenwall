package checks

import (
	"log"
	"reflect"

	"github.com/mtojek/greenwall/middleware/monitoring"
	"github.com/paulstuart/ping"
)

const (
	pingCheckName     = "ping_check"
	messageICMPFailed = "ICMP failed"

	pingTimeoutSeconds = 3
)

// PingCheck is ICMP health check.
// This check requires CAP_NET_RAW capability or to be run as root user.
type PingCheck struct {
	monitoringConfiguration *monitoring.Configuration
	nodeConfig              *monitoring.Node
}

// Initialize method initializes the check instance.
func (p *PingCheck) Initialize(monitoringConfiguration *monitoring.Configuration, nodeConfig *monitoring.Node) {
	p.monitoringConfiguration = monitoringConfiguration
	p.nodeConfig = nodeConfig
}

// Run method executes the check. This is invoked periodically.
func (p *PingCheck) Run() CheckResult {
	err := ping.Pinger(p.nodeConfig.Endpoint, pingTimeoutSeconds)
	if err != nil {
		log.Println(err)
		return CheckResult{
			Status:  StatusDanger,
			Message: err.Error(),
		}
	}
	return CheckResult{
		Message: MessageOK,
		Status:  StatusSuccess,
	}
}

func init() {
	registerType(pingCheckName, reflect.TypeOf(PingCheck{}))
}
