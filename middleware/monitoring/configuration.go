package monitoring

import "time"

// Configuration stores the configuration of monitored nodes.
type Configuration struct {
	General General `yaml:"general"`
	Groups  []Group `yaml:"groups"`
}

// General stores common settings of GreenWall.
type General struct {
	HealthcheckEvery      time.Duration `yaml:"healthcheckEvery"`
	HTTPClientTimeout     time.Duration `yaml:"hTTPClientTimeout"`
	RefreshDashboardEvery time.Duration `yaml:"refreshDashboardEvery"`
}

// Group stores definitions of monitored nodes.
type Group struct {
	Name  string `yaml:"name"`
	Nodes []Node `yaml:"nodes"`
}

// Node stores monitoring definition of a single node.
type Node struct {
	Name                    string `yaml:"name"`
	Endpoint                string `yaml:"endpoint"`
	HealthyResponseContains string `yaml:"healthyResponseContains"`
}
