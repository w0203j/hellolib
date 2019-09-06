package hellolib

import (
	"fmt"

	"rsc.io/quote"
)

// SayHello function
func SayHello(s string) {
	fmt.Printf("about to rsc.io/quote string %s", s)
	quote.Opt()
}
