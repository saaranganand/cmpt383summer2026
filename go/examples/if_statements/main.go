// if_statements/main.go

//
// This demonstrates the use of if statements in Go. They are fairly similar to
// those in C/C++, although:
//
// - The condition does not need to be enclosed in parentheses.
//
// - {}-braces are always required, even if the body is a single statement.
//
//

package main

import "fmt"

func main() {
	age := 0
	fmt.Print("How old are you? ")
	fmt.Scanln(&age)

	if age >= 65 {
		fmt.Println("You can retire!")
	}

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	if 13 <= age && age < 18 {
		fmt.Println("You are a teenager.")
	} else if 6 <= age && age < 13 {
		fmt.Println("You are a child.")
	} else if 0 <= age && age < 6 {
		fmt.Println("You are an infant.")
	} else {
		fmt.Println("You have very old or very young.")
	}

} // main
