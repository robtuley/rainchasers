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

func init() {
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
	logger := daemon.NewLogger("web")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/s/", http.StripPrefix("/s/", fs))

	http.HandleFunc("/", serveTemplate)

	logger.Info("server.start", report.Data{
		"port": port,
	})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Action("server.error", report.Data{
			"error": err.Error(),
		})
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	s, exists := sectionM[r.URL.Path]
	if exists {
		sectionT.ExecuteTemplate(w, "section", s)
		return
	}
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", 302)
		return
	}
	sectionT.ExecuteTemplate(w, "home", rainchasers.Sections)
}
