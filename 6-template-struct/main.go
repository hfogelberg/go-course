package main

import (
  "os"
  "log"
  "text/template"
)

var tpl  *template.Template

type quote struct {
  Name string
  Quote string
}

func init() {
  tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
  q := quote {
    Name: "Darth Vader",
    Quote: "I find your lack of faith disturbing",
  }
  err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", q)
  if err != nil {
    log.Fatalln(err)
  }
}
