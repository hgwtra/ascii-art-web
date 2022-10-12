package server

import (
	"html/template"
	"net/http"
)

//create a variable from server package
var Tpl *template.Template

func GetRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorNotFound(w)
		return
	}
	Tpl.Execute(w, nil)

}
