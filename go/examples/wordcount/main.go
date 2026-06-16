// wordcount/main.go

//
// Count the number of words in a file, and print a list of the top N most
// frequently occurring words. You can print the top N=10 words by default, but
// it should work with any value of N.
//
// For words with same count, sort them by alphabetical order in the list.
//
// Convert all uppercase letters A-Z to lowercase a-z, and replace characters
// that are not a-z with spaces.
//

package main

import (
	"fmt"
	"os"
	// "unicode"
	"regexp"
	"sort"
	"strings"
)

func main() {

	//
	// open the file
	//
	bytes, err := os.ReadFile("austenPandP.txt")
	if err != nil {
		panic("Couldn't open file!")
		// The idea of panic is that something so serious has happened that the
		// program can't continue. In this demo program, if we can't open the
		// file then there is no point in continuing.
	}

	//
	// convert letters to lowercase, non-letters to spaces
	//
	content := strings.ToLower(string(bytes))

	// rep is a regular expression that matches all characters that are *not*
	// one of the 26 lowercase letters a-z (the ^ means "not")
	rep := regexp.MustCompile(`[^a-z]`)

	// replace each non-letter with a space
	processedContent := rep.ReplaceAllString(content, " ")

	//
	// create a slice of all the words
	//
	words := strings.Split(processedContent, " ")

	//
	// trim any extra whitespace around each word
	//
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
	}

	//
	// create a map of the counts of all the words
	//
	freq := map[string]int{} // initially empty map
	for _, w := range words {
		if len(w) > 0 {
			freq[w] += 1
		}
	}

	//
	// create a key-value struct to for sorting the frequencies
	//
	type kv struct {
		key string
		val int
	}

	//
	// copy the frequencies into a slice of key-value pairs
	//
	data := []kv{} // initially empty slice of key-value pairs

	// add each key-value pair to the slice
	for k, v := range freq {
		item := kv{k, v}          // create a new key-value pair
		data = append(data, item) // add it to the slice
	}

	//
	// sort by alphabetically by word (the key)
	//
	// note that this is a function call of the form:
	//   sort.SliceStable(slice, function)
	//
	// a function literal is passed to sort.SliceStable to compare the items
	//
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].key < data[j].key
	})

	//
	// sort from most frequent to least frequent
	//
	// note that the comparison function is different than above
	//
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].val > data[j].val
	})

	//
	// print the top 10 most frequent words
	//
	for i, item := range data[:10] {
		fmt.Printf("%v. %v (%v)\n", i+1, item.key, item.val)
	}
} // main
