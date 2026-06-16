package main

import "fmt"

//
// Index returns the index of x in s, or -1 if not found. The first parameter is
// a slice of type T, the second is a value of type T.
//
// T is any type that supports the comparable constraint, i.e. T works with ==
// and !=.
//
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println("int slice Index: ", Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println("string slice Index: ", Index(ss, "hello"))
}
