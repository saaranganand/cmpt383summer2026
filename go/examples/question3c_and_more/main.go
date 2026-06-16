// question3c_and_more.go

//
// This demonstrates the use of closures in Go. A closure is function plus an
// environment of variables that the function can use.
//
// Go has good support for closures, and so passing and returning functions is a
// common pattern in Go.
//

package main

import "fmt"

func main() {
	// test_makeAdder()
	test_makeIncrementer()
}

//
// Returns a closure (a function with an associated environment of variables)
// that adds n to a given int. In other words, it makes a function that adds n
// to a number.
//
func makeAdder(n int) func(int) int {
	return func(a int) int { // lambda function
		return a + n
	}
}

func test_makeAdder() {
	add5 := makeAdder(5)
	n := 1
	fmt.Println("n is:", n)
	fmt.Printf("add5(%v) returns: %v\n", n, add5(n))
	fmt.Printf("add5(add5(%v)) returns: %v\n", n, add5(add5(n)))
}

//
// This returns two closures (functions with an associated environment of
// variables). 
//
// The first closure is an increment function that adds 1 to the variable n.
// Notice in the function header that the return type is fun(), i.e. a function
// that takes no input and returns nothing.
//
// The second closure is a getter function that returns the current value of n.
// Both closures share the same variable n, so it they act like an object with
// two methods: increment and get. Notice in the function header that the return
// type is func() int, i.e. a function that takes no input and returns an int.
//
func makeIncrementer() (func(), func() int) {
	// n is the shared value to be incremented
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
	fmt.Println("a has the value:", get_a())
	inc_a()
	fmt.Println("a has the value:", get_a())
	inc_a()
	inc_a()
	inc_a()
	fmt.Println("a has the value:", get_a())
}
