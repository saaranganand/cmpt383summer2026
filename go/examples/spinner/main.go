// spinner/main.go

//
// Draws an ASCII "spinner" graphic on the screen while doing a long
// calculation.
//
// From Chapter 8 of the book "The Go Programming Language", by Donovan and
// Kernighan.
//

package main

import (
	"fmt"
	"time"
)

func main() {
	const n = 45
	fmt.Printf("Calculating fib(%d) ...\n", n)

	// Launch a goroutine to spin the spinner. It will run in the background
	// while the main thread is calculating the Fibonacci number.
	go spinner(100 * time.Millisecond)

	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
