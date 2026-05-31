// numbered_list/main.go

package main

import "fmt"

func main() {
	pets := []string{"cat", "dog", "bird", "hamster"}
	for i, p := range pets {
		fmt.Printf("%v. %v\n", i+1, p)
	}
}
