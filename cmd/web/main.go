package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/rainchasers/content"
	"github.com/rainchasers/content/internal/river"
)

var sectionT *template.Template
var sectionM map[string]river.Section

func init() {
	f1 := filepath.Join("static", "section.html")
	f2 := filepath.Join("static", "badges.html")
	f3 := filepath.Join("static", "home.html")
	sectionT = template.Must(template.ParseFiles(f1, f2, f3))

	sectionM = make(map[string]river.Section, len(content.Sections))
	for _, s := range content.Sections {
		sectionM["/"+s.Slug] = s
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/s/", http.StripPrefix("/s/", fs))

	http.HandleFunc("/", serveTemplate)

	fmt.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	s, exists := sectionM[r.URL.Path]
	if exists {
		sectionT.ExecuteTemplate(w, "section", s)
		return
	}
	sectionT.ExecuteTemplate(w, "home", content.Sections)
}
