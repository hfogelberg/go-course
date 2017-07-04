package main

import (
  "os"
  "log"
  "text/template"
  "time"
)

var tpl *template.Template

func init( ) {
    tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func formatDate(t time.Time) string {
  return t.Format(time.RFC822)
}

var fm = template.FuncMap{
  "fdate": formatDate,
}

func main() {
  err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
  if err != nil {
    log.Fatalln(err)
  }
}
