package healthcheck

import (
	"regexp"
	"strings"

	"github.com/mtojek/greenwall/middleware/application"
	"github.com/mtojek/greenwall/middleware/monitoring"
)

const (
	statusDanger  = "danger"
	statusSuccess = "success"

	messageNotCheckedYet = "not checked yet"

	updateCheckChanLength  = 256
	anchorReplaceCharacter = "_"
)

var anchorRegexp = regexp.MustCompile("[^A-Za-z0-9]+")

// Healthcheck is responsible for storing latest statuses of monitored nodes.
type Healthcheck struct {
	applicationConfiguration *application.Configuration
	monitoringConfiguration  *monitoring.Configuration

	board HealthStatus

	requestStatusChan      chan bool
	requestUpdateCheckChan chan checkResult
	responseStatusChan     chan HealthStatus
}

type checkResult struct {
	groupOffset int
	nodeOffset  int
	status      string
	message     string
}

// NewHealthcheck method creates a new instance of healthcheck.
func NewHealthcheck(applicationConfiguration *application.Configuration,
	monitoringConfiguration *monitoring.Configuration) *Healthcheck {
	return &Healthcheck{
		applicationConfiguration: applicationConfiguration,
		monitoringConfiguration:  monitoringConfiguration,

		requestStatusChan:      make(chan bool),
		requestUpdateCheckChan: make(chan checkResult, updateCheckChanLength),
		responseStatusChan:     make(chan HealthStatus),
	}
}

// Start method starts the monitoring routines.
func (h *Healthcheck) Start() {
	h.fillBoard()
	go h.processRequests()
	// Run go routinow with sleep
	// http client, szukanie frazy i zapis wynikow
}

func (h *Healthcheck) fillBoard() {
	var groups []Group
	for _, configuredGroup := range h.monitoringConfiguration.Groups {
		var nodes []Node
		for _, configuredNode := range configuredGroup.Nodes {
			nodes = append(nodes, Node{
				Name:     configuredNode.Name,
				Endpoint: configuredNode.Endpoint,
				Status:   statusDanger,
				Message:  messageNotCheckedYet,
			})
		}
		groups = append(groups, Group{
			Name:   configuredGroup.Name,
			Anchor: h.asAnchor(configuredGroup.Name),
			Nodes:  nodes,
		})
	}
	h.board = HealthStatus{Groups: groups}
}

func (h *Healthcheck) asAnchor(name string) string {
	return strings.TrimSuffix(anchorRegexp.ReplaceAllString(name, anchorReplaceCharacter), anchorReplaceCharacter)
}

func (h *Healthcheck) processRequests() {
	for {
		select {
		case <-h.requestStatusChan:
			h.responseStatusChan <- h.copyOfBoard()
		case result := <-h.requestUpdateCheckChan:
			h.applyChange(result)
		}
	}
}

func (h *Healthcheck) copyOfBoard() HealthStatus {
	var copyOfGroups []Group

	for _, group := range h.board.Groups {
		var copyOfNodes []Node
		for _, node := range group.Nodes {
			copyOfNodes = append(copyOfNodes, Node{
				Name:     node.Name,
				Endpoint: node.Endpoint,
				Status:   node.Status,
				Message:  node.Message,
			})
		}
		copyOfGroups = append(copyOfGroups, Group{
			Name:   group.Name,
			Anchor: group.Anchor,
			Nodes:  copyOfNodes,
		})
	}
	return HealthStatus{Groups: copyOfGroups}
}

func (h *Healthcheck) applyChange(result checkResult) {
	h.board.Groups[result.groupOffset].Nodes[result.nodeOffset].Status = result.status
	h.board.Groups[result.groupOffset].Nodes[result.nodeOffset].Message = result.message
}

// Status method returns a report containing statuses of monitored nodes.
func (h *Healthcheck) Status() HealthStatus {
	h.requestStatusChan <- true
	return <-h.responseStatusChan
}

// UpdateBoard method stores new check result in the board.
func (h *Healthcheck) UpdateBoard(result checkResult) {
	h.requestUpdateCheckChan <- result
}
