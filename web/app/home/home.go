package home

import (
	"html/template"
	"net/http"
)

// Assuming you have a global template variable
var templates *template.Template

func init() {
	// Load your templates
	templates = template.Must(template.ParseGlob("web/template/*"))
}

// Handler for our home page.
func Handler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
