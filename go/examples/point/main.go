// point/main.go

package main

import (
	"fmt"
)

type Point struct {
	x, y float64
}

//
// String() below is a method. The difference between a method and a function is
// that a method has a receiver argument that appears before the method name.
//
// In this example, String() is a method with receiver p. p is of type Point,
// and is passed by value to the method.
//
// Any method with the signature "String() string" automatically implements the
// Stringer interface, which is pre-defined by Go:
//
//    type Stringer interface {
//      String() string
//    }
//
// Functions like Sprintf work with values that implement Stringer.
//
func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

//
// A good question to ask is this: why not write it like a regular function
// instead?
//
// func String(p Point) string {
//  return fmt.Sprintf("(%v, %v)", p.x, p.y)
// }
//
// The issue is that now the type of the function is String(Point) string. A
// function like fmt.Sprintf doesn't know about the Point type.
//
// But with the method version, the type is String() string for all methods, and
// so fmt.Sprintf can call it. The receiver is implicitly passed to the method.
//

func main() {
	p := Point{4.5, -0.44}
	fmt.Println(String(p))  // p.String()
	// q := Point{10, 100}
	// fmt.Println(p, q)
	// fmt.Printf("p = %v, q = %v\n", p, q)
}
