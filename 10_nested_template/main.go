package main

import (
  "os"
  "log"
  "text/template"
)

var tpl *template.Template

func init() {
  tpl = template.Must(ParseFiles("index.gohtml"))
}

func main() {

}
