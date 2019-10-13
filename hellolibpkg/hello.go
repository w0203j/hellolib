package hellolibpkg

import (
	"fmt"

	"rsc.io/quote"
)

// SayHello function
func SayHello(s string) {
	fmt.Printf("about to rsc.io/quote string %s", s)
	fmt.Println(quote.Opt())
}
