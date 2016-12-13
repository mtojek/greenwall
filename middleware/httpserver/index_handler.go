package httpserver

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/mtojek/greenwall/middleware/application"
	"github.com/mtojek/greenwall/middleware/healthcheck"
	"github.com/mtojek/greenwall/middleware/monitoring"
)

const indexFile = "index.html"

// IndexHandler is responsible for serving live dashboard page.
type IndexHandler struct {
	page                    *template.Template
	monitoringConfiguration *monitoring.Configuration
	healthcheck             *healthcheck.Healthcheck
}

// NewIndexHandler method creates a new instance of IndexHandler.
func NewIndexHandler(applicationConfiguration *application.Configuration,
	monitoringConfiguration *monitoring.Configuration, healthcheck *healthcheck.Healthcheck) *IndexHandler {
	page, err := template.New(indexFile).ParseFiles(path.Join(applicationConfiguration.StaticDir, indexFile))
	if err != nil {
		log.Fatalf("Error occurred while parsing template: %v", err)
	}
	return &IndexHandler{
		page: page,
		monitoringConfiguration: monitoringConfiguration,
		healthcheck:             healthcheck,
	}
}

// ServeHTTP method returns the live dashboard page (an applied Index template).
func (indexHandler *IndexHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	err := indexHandler.page.Execute(rw, indexHandler.readPageData())
	if err != nil {
		log.Printf("Error occurred while executing template: %v", err)
	}
}

func (indexHandler *IndexHandler) readPageData() *PageData {
	return &PageData{
		LastRefreshTime:       time.Now().Format(time.Stamp),
		RefreshDashboardEvery: indexHandler.monitoringConfiguration.General.RefreshDashboardEvery.Seconds(),
		HealthStatus:          indexHandler.healthcheck.Status(),
	}
}
