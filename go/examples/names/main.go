// names/main.go

//
// Given a list of names of the form "Last-name, First-name", convert them all
// to "First-name Last-name".
//

package main

import "fmt"

// Find the first location in s such that s[i] == c. Otherwise, returns -1.
func find(s string, c byte) int {
    for i := 0; i < len(s); i++ {
        if s[i] == c {
            return i
        }
    }
    return -1
}

func convert(s string) (string, string) {
    commaAt := find(s, ',')
    if commaAt == -1 { return s, "" }
    return s[commaAt+2:], s[:commaAt]
}

var names = []string{
    "Abbott, Bud", "Costello, Lou", "Laurel, Stan",
    "Hardy, Oliver", "Chaplin, Charlie", "Lloyd, Harold",
    "Keaton, Buster", "Marx, Groucho", "Marx, Harpo",
    "Sellers, Peter",
    }

func main() {
    for i, s := range names {
        first, last := convert(s)
        fmt.Printf("%v. %v %v\n", i+1, first, last)
    }
}