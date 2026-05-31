// functionsAndClosures.go

//
// This demonstrates the use of closures in Go. A closure is function plus an
// environment of variables that the function can use.
//
// Go has good support for closures, and so passing and returning functions is a
// common pattern in Go.
//

package main

import (
	"fmt"
)

func main() {
	test_makeAdder()
	// test_makeIncrementer()
}

//
// Returns a closure (a function with an associated environment of variables)
// that adds n to a given int. In other words, it makes a function that adds n
// to a number.
//
func makeAdder(n int) func(int) int {
	return func(a int) int {
		return a + n
	}
}

func test_makeAdder() {
	add5 := makeAdder(5)
	n := 1
	fmt.Printf("n=%v\n", n)
	fmt.Printf("n=%v\n", add5(n))
}

//
// An incrementer is a value that supports these two operations after it is
// created:
//
// - Add 1 (increment). - Return the current value
//
// makeIncrementer returns two closures, once for each of these operations. Both
// closes share the same variable n, and:
//
// - inc, the incrementer, takes no input and returns nothing; all you do is
//   call it and it adds 1 to n.
//
// - get, the getter, takes no input and returns the current value of n.
//
func makeIncrementer() (func(), func() int) {
	// n is the value to be incremented
	n := 0
	inc := func() {
		n++
	}
	get := func() int {
		return n
	}
	return inc, get
}

func test_makeIncrementer() {
	inc_a, get_a := makeIncrementer()
	inc_b, get_b := makeIncrementer()
	for i := 1; i <= 5; i++ {
		inc_a() // a is incremented once
		inc_b() // b is incremented twice
		inc_b()
	}
	fmt.Printf("get_a(): %v\n", get_a())
	fmt.Printf("get_b(): %v\n", get_b())
}
