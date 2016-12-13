package httpserver

import "github.com/mtojek/greenwall/middleware/healthcheck"

// PageData stores Index template data.
type PageData struct {
	LastRefreshTime       string
	RefreshDashboardEvery float64

	HealthStatus healthcheck.HealthStatus
}
