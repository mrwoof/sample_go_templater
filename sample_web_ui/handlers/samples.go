package handlers

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type SampleOne struct {
	BasePage
	Time time.Time
}

type SampleTwo struct {
	BasePage
	Input string
}

// GetSampleOne is the HTTP handler for the /sample_one path
func (h *Handlers) GetSampleOne(w http.ResponseWriter, r *http.Request) {
	templateData := &SampleOne{
		BasePage: BasePage{
			Title:    "Sample One",
			Username: "user1234",
		},
		Time: time.Now(),
	}

	log.Infof("serving sample_one.html with %#v", templateData)
	h.ExecuteTemplate(w, "sample_one.html", templateData)
}

// GetSampleTwo is the HTTP handler for the /sample_two path
func (h *Handlers) GetSampleTwo(w http.ResponseWriter, r *http.Request) {
	templateData := &SampleTwo{
		BasePage: BasePage{
			Title:    "Sample Two",
			Username: "user1234",
		},
		Input: r.FormValue("in"),
	}

	log.Infof("serving sample_two.html with %#v", templateData)
	h.ExecuteTemplate(w, "sample_two.html", templateData)
}
