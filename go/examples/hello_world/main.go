// hello_world/main.go

//
// This is the classic "Hello, world!" program in Go. Some things to note:
//
// - The package name is "main". Every Go program must have a main package where
//   the main function is defined.
//
// - The import statement imports the fmt package, which contains the Println
//   function. The quotes around fmt are required.
//
// - Functions start with the "func" keyword (so they are easy to search for).
//   The return type comes *after* the function name, and no function return type
//   is written if the function returns nothing.
//
// - Semi-colon ; characters are not required at the end of statements, and are
//   typically not written.
//
// - The format of the code is standard: the gofmt tool is generally use to
//   format Go code in a standard way.
//
// - To run the code type this at the command line:
//
//     $ go run main.go
//
//   This compiles and runs the program. One of Go's design goals is to make compiling fast,
//   and so this is a very fast way to run a program.
//
// - To create an executable binary file, type this at the command line:
//
//     $ go build main.go  # creates an executable file called main
//     $ ./main
//
//   The binary file size can be quite large, e.g. this hello world programis a 2.5MB file:
//
//     $ stat -f%z main
//     2508978
//
//   The executable files for Go tend to be big because they include the Go runtime,
//   e.g. a garbage collector, goroutine scheduler, support for reflection, etc.
//

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
