package views

import (
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	HTMLTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	err = t.HTMLTpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "There was an internal error executing the template.", http.StatusInternalServerError)
		return
	}
}
