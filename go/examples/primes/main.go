// primes/main.go

//
// This programs demonstrates various basic Go features combined together to
// generate prime numbers.
//

package main

import "fmt"

func main() {
    for i := 0; i <= 25; i++ {
        fmt.Println(i, IsPrime(i))
    }

    fmt.Println(PrimesLessThan(10000))  // 1229
}

// Returns true if the integer n is prime, and false otherwise.
func IsPrime(n int) bool {
    if n < 2 {
        return false
    } else if n == 2 {
        return true
    } else if n % 2 == 0 {
        return false
    } else {
        candidate := 3
        for candidate * candidate <= n {
            if n % candidate == 0 {
                return false
            }
            candidate += 2
        }
        return true
    }
}

// Uses a named return parameter.
func PrimesLessThan(n int) (result int) {
    for i := 2; i < n; i++ {
        if IsPrime(i) {
            result++
        }
    }
    return // required with named return parameter
}
