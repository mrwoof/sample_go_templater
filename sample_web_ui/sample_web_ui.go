package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"sample_app/sample_web_ui/handlers"
)

const serviceName = "sample_web_ui"

func main() {
	var (
		port         = 8080
		templatesDir = "../templates"
		staticDir    = "../static"
	)

	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		log.Fatalf("templates directory does not exist: %s", templatesDir)
	}
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("static directory does not exist: %s", staticDir)
	}

	h, err := handlers.NewHandlers(templatesDir)
	if err != nil {
		log.Fatalf("unable to init handlers: %s", err)
	}

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/sample_one", h.GetSampleOne)
	http.HandleFunc("/sample_two", h.GetSampleTwo)

	addr := fmt.Sprintf(":%d", port)
	log.Infof("%s started on 0.0.0.0%s", serviceName, addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("unable start server: %s", err)
	}
}
