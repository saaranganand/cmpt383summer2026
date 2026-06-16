// functions/main.go

//
// This program demonstrates the use of functions in Go. Notice a few things:
//
// - The type of the parameters comes *after* the parameter name, e.g. "x int"
//   instead of "int x". Note that alternative way to write the parameters as
//   "x, y int" is allowed.
//
// - The return type comes *after* the function name and before the function
//   body.
//
// - The swap function demonstrates the use of multiple return values, i.e. a
//   function that returns two values.
//
// - The avg_diff function demonstrates the use of **named return values**, i.e.
//   a function that returns two values with names avg and diff. The bare return
//   at the end is required. Note that the return types are named in the function
//   declaration.
//

package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func avg_diff(x, y float64) (avg float64, diff float64) {
	avg = float64(x+y) / 2
	diff = x - y
	return
}

func main() {
	sum := add(1, 2)
	fmt.Println("sum:", sum)

	a, b := "hello", "world"
	fmt.Println("before swap:", a, b)
	// a, b = swap(a, b)
	a, b = b, a
	fmt.Println("after swap:", a, b)

	avg, diff := avg_diff(1.0, 2.0)
	fmt.Println("avg:", avg, " diff:", diff)
}
