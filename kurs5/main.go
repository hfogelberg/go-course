package main

import (
  "os"
  "log"
  "text/template"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
  peeps :=map [string]string{"Luke": "Jeddi apprentice", "Han":"Smuggler", "Leia":"Princess", "Obi":"Jedi guru", "Darth":"Bad guy"}
  err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", peeps)
  if err != nil {
    log.Fatalln(err)
  }
}
