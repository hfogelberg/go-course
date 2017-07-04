package main

import (
  "os"
  "strings"
  "text/template"
  "log"
)

var tpl *template.Template

var fm = template.FuncMap{
  "uc": strings.ToUpper,
  "ft": firstThree,
}

type creature struct {
  Name string
  Being string
}

func init() {
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func firstThree(s string) string {
  s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
  creatures := []creature {
                  creature {"Dobbi", "Husalf"},
                  creature{"Fauks", "Fenix"},
                  creature{"Hagrid", "Halvjätte"},
                  creature{"Hedvig", "Kattuggla"},
                }

  err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", creatures)
  if err != nil {
    log.Fatalln(err)
  }
}
