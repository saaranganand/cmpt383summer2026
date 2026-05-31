// defer/main.go

// This demonstrates the use of defer. When you defer a statement the
// statement's is executed at the end of the function, but before the function
// returns.
//
// Defer is useful for things like closing files, unlocking mutexes, and
// releasing resources.
//

package main

import "fmt"

func main() {
	defer fmt.Println("All done!")
	fmt.Println("Doing some work...")
}
