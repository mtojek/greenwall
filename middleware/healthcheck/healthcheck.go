package healthcheck

import (
	"regexp"
	"strings"
	"time"

	"github.com/mtojek/greenwall/middleware/application"
	"github.com/mtojek/greenwall/middleware/healthcheck/checks"
	"github.com/mtojek/greenwall/middleware/monitoring"
)

const (
	updateCheckChanLength  = 256
	anchorReplaceCharacter = "_"
)

var anchorRegexp = regexp.MustCompile("[^A-Za-z0-9]+")

type checkResultInBoard struct {
	groupOffset int
	nodeOffset  int
	result      checks.CheckResult
}

// Healthcheck is responsible for storing latest statuses of monitored nodes.
type Healthcheck struct {
	applicationConfiguration *application.Configuration
	monitoringConfiguration  *monitoring.Configuration

	board HealthStatus

	requestStatusChan      chan bool
	requestUpdateCheckChan chan checkResultInBoard
	responseStatusChan     chan HealthStatus
}

// NewHealthcheck method creates a new instance of healthcheck.
func NewHealthcheck(applicationConfiguration *application.Configuration,
	monitoringConfiguration *monitoring.Configuration) *Healthcheck {
	return &Healthcheck{
		applicationConfiguration: applicationConfiguration,
		monitoringConfiguration:  monitoringConfiguration,

		requestStatusChan:      make(chan bool),
		requestUpdateCheckChan: make(chan checkResultInBoard, updateCheckChanLength),
		responseStatusChan:     make(chan HealthStatus),
	}
}

// Start method starts the monitoring routines.
func (h *Healthcheck) Start() {
	h.fillBoard()
	go h.processRequests()
	h.runChecks()
}

func (h *Healthcheck) fillBoard() {
	var groups []Group
	for _, configuredGroup := range h.monitoringConfiguration.Groups {
		var nodes []Node
		for _, configuredNode := range configuredGroup.Nodes {
			var endpoint string
			if len(configuredNode.Endpoint) > 0 {
				endpoint = configuredNode.Endpoint
			} else {
				endpoint = ""
			}

			nodes = append(nodes, Node{
				Name:         configuredNode.Name,
				Type:         configuredNode.Type,
				Endpoint:     endpoint,
				HTTPEndpoint: strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://"),
				Status:       checks.StatusWarning,
				Message:      checks.MessageNotCheckedYet,
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

func (h *Healthcheck) runChecks() {
	for i, group := range h.monitoringConfiguration.Groups {
		for j, node := range group.Nodes {
			go h.runCheck(i, j, node)
		}
	}
}

func (h *Healthcheck) runCheck(groupOffset, nodeOffset int, nodeConfig monitoring.Node) {
	routine := checks.MakeInstance(nodeConfig.Type)
	routine.Initialize(h.monitoringConfiguration, &nodeConfig)

	for {
		time.Sleep(h.monitoringConfiguration.General.HealthcheckEvery)
		resultInBoard := checkResultInBoard{
			groupOffset: groupOffset,
			nodeOffset:  nodeOffset,
			result: checks.CheckResult{
				Status: checks.StatusDanger,
			},
		}
		resultInBoard.result = routine.Run()
		h.UpdateBoard(resultInBoard)
	}
}

func (h *Healthcheck) copyOfBoard() HealthStatus {
	var copyOfGroups []Group

	for _, group := range h.board.Groups {
		var copyOfNodes []Node
		for _, node := range group.Nodes {
			copyOfNodes = append(copyOfNodes, Node{
				Name:         node.Name,
				Type:         node.Type,
				Endpoint:     node.Endpoint,
				HTTPEndpoint: node.HTTPEndpoint,
				Status:       node.Status,
				Message:      node.Message,
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

func (h *Healthcheck) applyChange(resultInBoard checkResultInBoard) {
	h.board.Groups[resultInBoard.groupOffset].Nodes[resultInBoard.nodeOffset].Status = resultInBoard.result.Status
	h.board.Groups[resultInBoard.groupOffset].Nodes[resultInBoard.nodeOffset].Message = resultInBoard.result.Message
}

// Status method returns a report containing statuses of monitored nodes.
func (h *Healthcheck) Status() HealthStatus {
	h.requestStatusChan <- true
	return <-h.responseStatusChan
}

// UpdateBoard method stores new check result in the board.
func (h *Healthcheck) UpdateBoard(result checkResultInBoard) {
	h.requestUpdateCheckChan <- result
}
