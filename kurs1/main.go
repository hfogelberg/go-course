package main

import (
  "fmt"
)

type person struct {
            fname string
            lname string
          }

  type jedi struct {
            person
            hasForce bool
  }

  func (p person) speak() {
    fmt.Println("I want money, says " + p.fname + " " + p.lname)
  }

  func(j jedi) speak() {
    fmt.Println(j.fname + " says May the force be with you")
  }

  type human interface {
    speak()
  }

func main() {
  p1 := person{"Han", "Solo"}
  fmt.Println(p1)

  j1 := jedi{
    person {
                "Luke",
              "Skywalker",
            } ,
            true,
        }
  p1.speak()
  j1.speak()
}
