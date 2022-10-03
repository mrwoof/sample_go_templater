package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"sample_app/internal/templater"
)

type Handlers struct {
	Templater *templater.Templater
}

type BasePage struct {
	Title    string
	Username string
}

// NewHandlers creates a Handlers object and loads the HTML templates for it.
func NewHandlers(templateDir string) (*Handlers, error) {
	t, err := templater.LoadTemplates(templateDir)
	if err != nil {
		return nil, err
	}

	return &Handlers{
		Templater: t,
	}, nil
}

// ExecuteTemplate calls Execute for the given matching template name using the data passed in
func (h *Handlers) ExecuteTemplate(w http.ResponseWriter, name string, d interface{}) error {
	err := h.Templater.ExecuteTemplate(w, name, d)
	if err != nil {
		log.WithError(err).Errorf("failed rendering template %q", name)
		http.Error(w, "error rendering template", 500)
	}

	return err
}
