// question5/main.go

package main

import "fmt"

//
// An answer to Question 5 of the Go problem set:
//
// Write a Go program that uses each of the following functions. In the
// comments, give clear descriptions of what each function does:
//
// - make([]T, len, cap)
// - append(s, elem)
// - append(s, elems...)
// - copy(dst, src)
// - len(s)
// - cap(s)
// - clear(s) (sets the elements of slice s to the zero value; length remains
//   unchanged)
//

// global indent string constant
const indent = "   "

// print a message on its own line
func example(msg string) {
	fmt.Printf("\n%v:\n", msg)
}

func printNamedSliceInfo(name string, s []int) {
	fmt.Printf("%v%v: %v len=%v cap=%v\n", indent, name, s, len(s), cap(s))
}

func main() {
	//
	// make([]T, len, cap) creates a new slice of type T with length len and capacity cap.
	// The values of the slice are all initialized to 0.
	//
	s := make([]int, 5, 10)
	example("after make([]int, 5, 10)")
	printNamedSliceInfo("s", s)

	//
	// set the elements of the slice to 0, 1, 2, 3, 4
	//
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	example("after setting the elements C-style for-loop")
	printNamedSliceInfo("s", s)

	//
	// add 10 to each element of the slice
	//
	for i := range s {
		s[i] += 10
	}
	example("after adding 10 to each element")
	printNamedSliceInfo("s", s)

	//
	// clear(s) sets the elements of slice s to the zero value; length remains unchanged.
	//
	clear(s)
	example("after clear(s)")
	printNamedSliceInfo("s", s)

	//
	// set s to a new empty slice
	//
	s = make([]int, 3)
	example("after setting s to make([]int, 3)")
	printNamedSliceInfo("s", s)

	//
	// append 1, 2, 3
	//
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	example("after appending 1, 2, 3 to s")
	printNamedSliceInfo("s", s)

	//
	// append 4, 5, 6
	//
	s = append(s, 4, 5, 6)
	example("after appending 4, 5, 6 to s")
	printNamedSliceInfo("s", s)

	//
	// make a new slice {1, 2, 3}
	//
	s = []int{1, 2, 3}
	example("after setting s to []int{1, 2, 3}")
	printNamedSliceInfo("s", s)

	//
	// append s to s
	//
	s = append(s, s...)
	example("after appending s to s")
	printNamedSliceInfo("s", s)

	//
	// append a new slice t to s
	//
	t := []int{4, 5, 6}
	s = append(s, t...)
	example("after appending t to s")
	printNamedSliceInfo("s", s)

	//
	// copy example 1
	//
	example("copy example 1")
	src := []int{1, 2, 3}
	dst := make([]int, 5)

	example("before copy")
	printNamedSliceInfo("src", src)
	printNamedSliceInfo("dst", dst)

	copy(dst, src)
	example("after copy(dst, src)")
	printNamedSliceInfo("src", src)
	printNamedSliceInfo("dst", dst)

	// //
	// // copy example 2
	// //
	// example("copy example 2")
	// src = []int{1, 2, 3}
	// dst = make([]int, 5)

	// example("before copy")
	// printNamedSliceInfo("src", src)
	// printNamedSliceInfo("dst", dst)

	// copy(dst[3:], src)
	// example("after copy(dst[2:], src)")
	// printNamedSliceInfo("src", src)
	// printNamedSliceInfo("dst", dst)

	// //
	// // copy example 3
	// //
	// example("copy example 3")
	// s = []int{1, 2, 3, 0, 0, 0}

	// example("before copy")
	// printNamedSliceInfo("s", s)

	// copy(s[1:], s)
	// example("after copy(s[1:], s)")
	// printNamedSliceInfo("s", s)

} // main
