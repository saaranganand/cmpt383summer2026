// age/main.go

// This demonstrates the use of the := operator to declare and initialize
// variables.
//
// The statement "name := "Fred"" declares a variable name and initializes it to
// the string "Fred". Go automatically infers the type of the variable, i.e.
// name is of type string.
//
// Similar for the statement "age := 42": it declares a variable age and
// initializes it to
//
// The alternate mains show equivalent ways to declare and initialize variables.
//

package main

import "fmt"

func main() {
	name := "Fred"
	age := 42
	fmt.Printf("%v is %v years old.\n", name, age)
}

// func main() {
// 	var name string = "Fred"
// 	var age int = 42
// 	fmt.Printf("%v is %v years old.\n", name, age)
// }

// func main() {
// 	var name string
// 	name = "Fred"
// 	var age int
// 	age = 42
// 	fmt.Printf("%v is %v years old.\n", name, age)
// }
