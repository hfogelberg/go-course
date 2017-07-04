package main

import (
  "fmt"
  "bufio"
  "strings"
)

func main() {
  s := "So, so you think you can tell\nHeaven from hell\nBlue skies from pain\nCan you tela green field\nFrom a cold steel rail?\nA smile from a veil?\nDo you think you can tell?\nDid they get you to trade\nYour heroes for ghosts?\nHot ashes for trees?\nHot air for a cool breeze?\nCold comfort for change?\nAnd did you exchange\nA walk on part in the war\nFor a lead role in a cage?\n"

  scanner := bufio.NewScanner(strings.NewReader(s))
  for scanner.Scan(){
    line := scanner.Text()
    fmt.Println(line)
  }
}
