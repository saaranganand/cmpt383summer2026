// splitter/main.go

//
// Gets the next word of a string (using spaces as delimiters), returns it as a
// string, and also returns the remaining string.
//

package main

import (
	"fmt"
)

func nextWord(s string) (string, string) {
	i := 0
	// skip leading spaces
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// if the string is empty, return empty strings
	if i == len(s) {
		return "", ""
	}
	// find the end of the word
	j := i
	for j < len(s) && s[j] != ' ' {
		j++
	}
	return s[i:j], s[j:]
}

func words(s string) []string {
	words := []string{} 
	for s != "" {
		w, remaining := nextWord(s)
		words = append(words, w)
		s = remaining
	}
	return words
}

func quote(s string) string {
	return "\"" + s + "\""
}

func main() {
	s := "  Hello, world! "
	word, remaining := nextWord(s)
	fmt.Println("     word:", quote(word))
	fmt.Println("remaining:", quote(remaining))
	fmt.Println()

	word, remaining = nextWord(remaining)
	fmt.Println("     word:", quote(word))
	fmt.Println("remaining:", quote(remaining))
	fmt.Println()

	fmt.Println("test of words:")
	words := words(s)
	fmt.Println("words:", words)
}
