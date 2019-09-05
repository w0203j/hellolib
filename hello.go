package hellolib

import (
  "fmt"
  "rsc.io/quote"
)

func SayHello(string s) {
  fmt.Println("about to rsc.io/quote string %s", s)
  quote.Opt()
}
