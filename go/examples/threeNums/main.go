// threeNums/main.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Returns true if y is between x and z.
func oneTrial() bool {
	x := rand.Float64()
	y := rand.Float64()
	z := rand.Float64()
	return (x < y && y < z) ||  (z < y && y < x)
}

func main() {
	// seed the random number generator with a random value based on the
	// current time
	rand.Seed(time.Now().UnixNano())

	// Repeat the trial many times, keeping a count of how many times it
	// succeeds.
	numTrials := 100000
	count := 0
	for i := 0; i < numTrials; i++ {
		if oneTrial() {
			count++
		}		
	}

	// should be 1/3
	fmt.Printf("After %v trials: %v/%v = %v\n",
				numTrials, count, numTrials,
			    float64(count) / float64(numTrials))

} // main
