package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/rainchasers/content"
)

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
	fp := filepath.Join("static", "section.html")
	tmpl := template.Must(template.ParseFiles(fp))

	tmpl.ExecuteTemplate(w, "section", content.Sections[0])
}
