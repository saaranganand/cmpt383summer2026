// count_down/main.go

//
// This programs counts down using for-loop similar to the count_up program.
//

package main

import "fmt"

func main() {
	fmt.Println("Launch in")
	for i := 10; i > 0; i-- {
		fmt.Printf("... %v\n", i)
	}
	fmt.Println("Blast off!")
}
