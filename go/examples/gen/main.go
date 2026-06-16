// gen/main.go

//
// Demonstrates how to uses goroutines and channels to make generators (similar
// to Python generators).
//
// Intuitively, you can think of the channel like a little server in the
// computers memory that gives a value every time you ask for it using the <-
// operator.
//

package main

import "fmt"

// returns an int channel
func counter(n int) chan int {
	ch := make(chan int)

	// Launch a goroutine.
	go func() {
		for i := 0; i < n; i++ {
			ch <- i // blocks here until it is removed from the channel
		}
		close(ch)
	}() // note the () here: this is a function call

	// return the channel that communicates with the goroutine
	return ch
}

func TestCounter() {
	for n := range counter(10) {
		fmt.Println(n)
	}
}

// generate the Fibonacci numbers one at a time
func fibgen() chan int {
	ch := make(chan int)
	go func() {
		a, b := 1, 1
		for { // infinite loop
			ch <- a // blocks here until a is removed from channel
			a, b = b, a+b
		}
	}() // note the () here: this is a function call
	return ch
}

func TestFib() {
	// every time <-nextFib is called, the next Fibonacci number is taken from
	// the channel
	nextFib := fibgen()
	// for i := 0; i < 10; i++ {
	for range 10 {
		fmt.Println(<-nextFib)
	}
}

func main() {
	// TestCounter()
	TestFib()
}
