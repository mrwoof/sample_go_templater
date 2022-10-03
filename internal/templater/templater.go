package templater

import (
	"fmt"
	html "html/template"
	"net/http"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type Templater struct {
	Templates map[string]*html.Template
}

// LoadTemplates loads HTML templates into a map for reference by name of html file.
// Any template can be rendered (executed) by doing the following:
//	templateMap, err := LoadTemplates("../templates")
//	...
//	templateData := ... // a struct that holds data for rendering the template.
//	templateMap["my_page.html"].Exectue(w, templateData)
func LoadTemplates(templateDir string) (*Templater, error) {
	templates := make(map[string]*html.Template)

	// globalFile is an html file that is typically used by every individual template
	globalFile := filepath.Join(templateDir, "base.html")

	files, err := filepath.Glob(filepath.Join(templateDir, "*.html"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		name := filepath.Base(file)
		t := html.New(name)

		// add global functions to template prior to parsing
		t.Funcs(FuncMap)

		_, err := t.ParseFiles(file, globalFile)
		if err != nil {
			return nil, err
		}

		templates[name] = t
	}

	log.Debugf("html templates loaded: %#v", templates)

	return &Templater{Templates: templates}, nil
}

// ExecuteTemplate calls template.Execute for the given matching template name
// using the data passed in.
// An error is returned if no template has been loaded with that name.
func (templater *Templater) ExecuteTemplate(w http.ResponseWriter, name string, d interface{}) error {
	t, ok := templater.Templates[name]
	if !ok {
		return fmt.Errorf("template not found: %q", name)
	}

	return t.Execute(w, d)
}
