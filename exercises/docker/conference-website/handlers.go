package main

import (
	"html/template"
	"net/http"
)

var (
	templates = template.Must(template.New("").
		ParseGlob("templates/*.html"))
)

func (cw *conferenceWebsiteServer) homeHandler(w http.ResponseWriter, r *http.Request) {

	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"conference_name":   "DevNexus", 
	}); err != nil {
        
    }
}
