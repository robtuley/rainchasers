package main

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/robtuley/rainchasers"
	"github.com/robtuley/rainchasers/internal/daemon"
	"github.com/robtuley/rainchasers/internal/river"
	"github.com/robtuley/report"
)

const port = ":8080"

var sectionT *template.Template
var sectionM map[string]river.Section
var logger *report.Logger
var version string

type homePage struct {
	Version  string
	Sections []river.Section
}

func init() {
	logger = daemon.NewLogger("web")

	f1 := filepath.Join("static", "section.html")
	f2 := filepath.Join("static", "badges.html")
	f3 := filepath.Join("static", "home.html")
	sectionT = template.Must(template.ParseFiles(f1, f2, f3))

	sectionM = make(map[string]river.Section, len(rainchasers.Sections))
	for _, s := range rainchasers.Sections {
		sectionM["/"+s.Slug] = s
	}
}

func main() {
	// setup routes
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/s/", http.StripPrefix("/s/", fs))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", serveTemplate)

	// start server
	logger.Info("http.start", report.Data{
		"version": version,
		"port":    port,
	})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Action("http.stop", report.Data{
			"version": version,
			"error":   err.Error(),
		})
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	// try to serve a section
	s, exists := sectionM[r.URL.Path]
	if exists {
		logger.Info("http.response", report.Data{
			"status":  200,
			"path":    r.URL.Path,
			"section": s.UUID,
		})
		sectionT.ExecuteTemplate(w, "section", s)
		return
	}

	// redirect to home
	if r.URL.Path != "/" {
		logger.Info("http.response", report.Data{
			"status": 302,
			"path":   r.URL.Path,
		})
		http.Redirect(w, r, "/", 302)
		return
	}

	// serve homepage
	logger.Info("http.response", report.Data{
		"status": 200,
		"path":   r.URL.Path,
	})
	sectionT.ExecuteTemplate(w, "home", homePage{
		Version:  version,
		Sections: rainchasers.Sections,
	})
}
