package main

import (
  "os"
  "log"
  "text/template"
)

var tpl *template.Template

type creature struct {
  Name string
  Being string
}

type quote struct {
  Name string
  Says string
}

type stuff struct {
  Creatures []creature
  Quotes []quote
}

func init() {
    tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
  creatures := []creature {
                  creature {"Dobbi", "Husalf"},
                  creature{"Faulks", "Fenix"},
                  creature{"Hagrid", "Halvjätte"},
                  creature{"Hedvig", "Kattuggla"},
                }

    quotes := []quote {
                quote {"Luke", "May the force be with you"},
                quote {"Han Solo", "I want money"},
                quote {"Darth Vader", "I find your lack of faith disturbing"},
              }

    items := stuff {
                              Creatures: creatures,
                              Quotes: quotes,
                            }

    err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", items)
    if err != nil {
      log.Fatalln(err)
    }
}
