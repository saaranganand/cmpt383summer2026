// hello_name/main.go

//
// This program reads the user's name from the console and greets them. Notice a
// few things:
//
// - The variable name is declared with the var keyword. The type comes after
//   the variable name --- which seems strange to C/C++ programmers.
//
// - When name is created is given the empty string as its default value. In
//   general, Go assigns sensible default values to variables when they are
//   created.
//
// - The Scanf function reads the user's name from the console. Scanf is similar
//   to scanf in C/C++. Note that function parameters in Go are passed by value,
//   i.e. a copy of the variable's value is passed to the function. Since Scanf
//   modifies the value of name, it must be passed by reference, i.e. &name is
//   the address of name.
//
// - Go's string type is high-level and easy to use, e.g. the + operator is used
//   to concatenate strings.
//

package main

import "fmt"

func main() {
    fmt.Print("What's your name? ")
    var name string
    fmt.Scanf("%s", &name)
    fmt.Println("Good day " + name + ", how are you?")
}
