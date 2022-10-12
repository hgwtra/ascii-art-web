package main

import (
	"ascii/server"
	"html/template"

	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	//read the template
	tmpl = template.Must(template.ParseFiles("template/index.html"))
}

func main() {
	// Http. handle(pattern, handler Http.handler) => http.Fileserver(root) return handler => dir : root directory
	// call variable from server package and add value
	server.Tpl = tmpl
	//handle css from static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//get request
	http.HandleFunc("/", server.GetRequest)
	//post request
	http.HandleFunc("/ascii-art", server.PostRequest)
	//open port- listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
