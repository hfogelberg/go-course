package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type submitForm bool

func (s submitForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("My-Key", "This is Henrik's key")
	w.Header().Set("Content-Type", "text/html;charset:utf-8")
	fmt.Fprintln(w, "<h1>Hello Response</h1>")
}

func main() {
	var s submitForm
	http.ListenAndServe(":8080", s)
}
