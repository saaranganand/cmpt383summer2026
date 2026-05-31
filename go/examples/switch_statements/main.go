// switch_statements/main.go

//
// This demonstrates the use of switch statements in Go. They are fairly similar
// to those in C/C++, although:
//
// - The switch condition does not need to be enclosed in parentheses.
//
// - Switching can occur on any type, not just integral types. In the example
//   below it's switching on strings, which is not possible in C/C++.
//
// - Switch cases do *not* fall through by default. Go's switch runs the
//   statements just in the case that matches the condition, and does *not*
//   automatically fall through to the next case. Thus, a break statement isn't
//   required.
//
// - The alternate main show a switch statement with no condition: it is
//   equivalent to series of if-statements.
//

package main

import (
	"fmt"
	"runtime"
	// "time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening.")
// 	}
// }
