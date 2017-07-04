package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type webpage bool

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func (wp webpage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Handle form")
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submission    url.Values
		Header        http.Header
		ContentLength int64
		
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.html", data)
}

func main() {
	var wp webpage
	http.ListenAndServe(":8080", wp)
}
