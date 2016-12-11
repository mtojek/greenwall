package httpserver

import (
	"net/http"

	"html/template"
	"log"
	"path"

	"time"

	"github.com/mtojek/greenwall/middleware/configuration"
)

const indexFile = "index.html"

// IndexHandler is responsible for serving live dashboard page.
type IndexHandler struct {
	page *template.Template
}

// NewIndexHandler method creates a new instance of IndexHandler.
func NewIndexHandler(applicationConfiguration *configuration.ApplicationConfiguration) *IndexHandler {
	page, err := template.New(indexFile).ParseFiles(path.Join(applicationConfiguration.StaticDir, indexFile))
	if err != nil {
		log.Fatalf("Error occurred while parsing template: %v", err)
	}

	return &IndexHandler{
		page: page,
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
		TimeNow: time.Now().Format(time.RFC1123Z),
	}
}
