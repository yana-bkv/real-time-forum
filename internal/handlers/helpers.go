package handlers

import (
	"log"
	"net/http"
	"text/template"
)

// Main page
func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate("./web/html/index.html", w, nil)
}

// Code for rendering specific html file (repeated part)
func RenderTemplate(fileName string, w http.ResponseWriter, data interface{}) {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		http.Error(w, "Error: template parsing", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error: template executing", http.StatusInternalServerError)
		log.Println(err)
	}
}
