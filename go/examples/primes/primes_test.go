// prime_test.go

/*

This file shows how you can use Go's built-on test framework to implement
table-based testing.

To run this, in the same directory type this command:

   $ go test 
   PASS ok      
   examples/primes 0.180s

The "go test" automatically looks for files then with _test.go, and runs their
test functions.

A test function name must start with Test, and take a *testing.T as a parameter.
The test function then uses this pointer to signal failure and other such things 
to the test runner.

*/

package main

import "testing"

// Go test function names always begin with "Test" and take *testing.T as a
// parameter.
func TestIsPrime(t *testing.T) {
	// table-base testing
	some_primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 37, 101}
	some_non_primes := []int{-2, -1, 0, 1, 4, 6, 8, 9, 10, 25, 1001}

	// check primes
	for _, p := range some_primes {
		result := IsPrime(p)
		if result == false {
			// t.Errorf tells the Go test runner that this test has failed, and
			// gives an error message to print
			t.Errorf("IsPrime(%v) = false, should be true", p)
		}
	}

	// check non-primes
	for _, n := range some_non_primes {
		result := IsPrime(n)
		if result == true {
			t.Errorf("IsPrime(%v) = true, should be false", n)
		}
	}
} // TestIsPrime

// Go test function names always begin with "Test" and take *testing.T as a
// parameter.
func TestPrimesLessThan(t *testing.T) {
	// tests is a slice of structs, where each struct has an int n (the input to
	// PrimesLessThan), and int count (the expected output for n).
	tests := []struct {
		n     int
		count int
	}{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 3},
		{99, 25},
		{100, 25},
		{101, 25},
		{1000, 168},
		{10000, 1229},
	}

	for _, tc := range tests {
		result := PrimesLessThan(tc.n)
		if result != tc.count {
			// t.Errorf tells the Go test runner that this test has failed, and
			// gives an error message to print
			t.Errorf("PrimesLessThan(%v) = %v, should be %v", tc.n, result, tc.count)
		}
	}
} // TestPrimesLessThan
