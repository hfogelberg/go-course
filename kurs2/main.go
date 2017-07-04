package main

import (
  "log"
  "os"
  "text/template"
)

func main() {
  tpl, err := template.ParseFiles("index.gohtml")
  if err != nil {
		log.Fatalln(err)
	}

  // newFile, err := os.Create("index.html")
  // if err != nil {
  //   log.Println("Error creating new file")
  // }
  // defer newFile.Close()

  err = tpl.Execute(os.Stdout, nil)
  if err != nil {
		log.Fatalln(err)
	}
}
