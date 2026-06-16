// slices/main.go

//
// A slice is a "fat" pointer, i.e. it is a data structure that refers to a
// contiguous slice of an underlying array.
//
// In most cases, slices are preferred over arrays. But be careful: different
// slices might be backed by the same underlying array, and so modifying one
// slice might modify the other slice.
//

package main

import (
	"fmt"
)

func main() {
	//
	// arr is an array of 5 ints
	//
	// A Go array is like a C array: it is a contiguous block of memory, and it
	// can't be made bigger or smaller. Furthermore, the type of the array
	// includes the size of the array, e.g. arr is of type [5]int, not just
	// int[]. The 5 is part of the type.
	//
	arr := [5]int{1, 2, 3, 4, 5}

	//
	// A ranged for-loop is similar to Python's for-loop with enumerate(). The
	// first value is the index, and the second value is the value at that
	// index.
	//
	for i, val := range arr {
		fmt.Printf("arr[%v]: %v\n", i, val)
	}
	fmt.Println()

	//
	// A slice is a fat pointer: it refers to a part of an array without making
	// a copy of the values.
	//
	// In the code below, sl is a slice that refers to the elements arr[1],
	// arr[2], and arr[3].
	//
	// In general, arr[start:end] is a slice that refers to arr[start],
	// arr[start+1], ..., arr[end-1].
	//
	// Go's slice notation is similar to Python's slice notation.
	//
	var sl []int = arr[1:4]
	for x := range sl {
		fmt.Printf("sl[%v]: %v\n", x, sl[x])
	}
	fmt.Println()

	//
	// You can make slice from a slice. sl2 refers to the same underlying array
	// as sl.
	//
	sl2 := sl[1:2]
	for x := range sl2 {
		fmt.Printf("sl2[%v]: %v\n", x, sl2[x])
	}
	fmt.Println()

	fmt.Println("incAll(sl) ...")
	incAll(sl)
	
	// sl is modified
	for i, val := range sl {
		fmt.Printf("sl[%v]: %v\n", i, val)
	}
	fmt.Println()

	// print the original array
	for i, val := range arr {
		fmt.Printf("arr[%v]: %v\n", i, val)
	}
	fmt.Println()

	// you can use make to create a slice directly, e.g. s is a slice of ints
	// with length 10 and capacity 15
	s := make([]int, 10, 15)
	// s := []int{5, 3, 7, 8, 9, 5}
	fmt.Println("Length of s:", len(s))
	fmt.Println("Capacity of s:", cap(s))
	for i := range s {
		s[i] = i
	}
	for i, val := range s {
		fmt.Printf("s[%v]: %v\n", i, val)
	}
	fmt.Println()
} // main

// When you pass a slice to a function, the slice is copied but no copy of the
// underlying array is made, i.e. it acts as if the array were passed by
// reference.
func incAll(s []int) {
	// with only one range variable, the range variable is the index (not the value)
	for i := range s {
		s[i]++
	}
}
