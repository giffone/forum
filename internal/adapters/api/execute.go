package api

import (
	"forum/internal/constant"
	"html/template"
	"log"
	"net/http"
)

func Parse(path ...string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(path...)
	if err != nil {
		log.Printf("parseFile: %v", err)
		return nil, err
	}
	return tmpl, nil
}

func Execute(w http.ResponseWriter, tmpl *template.Template, define string, data interface{}) {
	log.Println("execute in")
	err := tmpl.ExecuteTemplate(w, define, data)
	if err != nil {
		log.Println("execute in err")
		log.Printf("executeTemplate: %v", err)
		ErrorsHTTP(w, "", constant.Code500)
	}
}
