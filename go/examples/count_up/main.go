// count_up/main.go

//
// This program counts up from 1 to 10 and prints each number. Notice a few
// things:
//
// - The for loop is similar to the for loop in C/C++. But no round parentheses
//   are used, and the {} around the body is required even if the body is a single
//   statement.
//
// - The i++ syntax is used to increment the loop variable i. This is similar to
//   the i++ syntax in C/C++. ++i is not allowed in Go.
//

package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Printf("This is sentence %v.\n", i)
    }
}
