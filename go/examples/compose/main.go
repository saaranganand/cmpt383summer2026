// functionsAndClosures.go

//
// This demonstrates the use of closures in Go to make a string compose
// function.
//
// compose(f, g) takes to string functions as input, and returns a new string
// function that first applies g to its input, and then applies f to the result.
// compose is an example of a higher-order function, and it returns a closure.
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Testing compose ...")
	makeTitle := compose(underline, strings.TrimSpace)
	s := "  This is a test! "
	title := makeTitle(s)
	fmt.Println(title)

}

// to cut down on clutter, we create a new type that describes the signature of
// string functions: any function that takes one strings as input, and returns
// one string as output.
type strFunc func(string) string

// Returns a new function h such that h(s) = f(g(s)).
func compose(f, g strFunc) strFunc {
	return func(s string) string {
		return f(g(s))
	}
}

// underline(s) returns a new string that is the same as s, but with a line of
// dashes underneath it
func underline(s string) string {
	bar := strings.Repeat("-", len(s))
	return s + "\n" + bar
}
