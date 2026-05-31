# Go Problem Set

The questions on the Go quiz will mainly be variations of the questions below,
or questions that are similar.

Please post your answers to the discussion board to share with other students.

**Important**: *Treat these problem-sets as non-AI activities!* Turn off all AI
support and try to figure them out yourself. Having AI or another student do
this for you will not help you learn. You must do the learning yourself!

## Question 1

Show three different ways in Go to create a new variable initialized to the
value `"cat"`.

## Question 2

Answer True or False for each of these questions about Go:

a. The `:=` operator is used to declare and initialize a variable without giving
   an explicit type.

b. The `switch` statement does NOT need a condition.

c. The `switch` statement automatically falls through to the next case (as in
   C/C++).

d. There are **no** while-loops in Go (that use `while` as the keyword).

e. Slices and arrays are the same thing.

f. You can make a slice from a slice.

g. The built-in `make` function can create a new slice.

h. Functions can return multiple values.

i. Function parameters are passed by value.

## Question 3

a. Write a Go function called `min_max` that takes two numbers as input, and
   returns a pair of values: the first value is the min of the two inputs, and
   the second value is the max.

   Use named return values for the function.

  It should work like this:

   ```go
   lo, hi := min_max(7, 5)
   fmt.Println(lo, hi) // prints 5 7

   lo, hi = min_max(1, 4)
   fmt.Println(lo, hi) // prints 1 4
   ```

b. Write a Go function called `min_and_max` that takes a slice of floating point
   numbers as input, and returns the min and max values in the slice. You can
   assume the slice is not empty.

   For example:

   ```go
   nums := []float64{3.1, 2.7, 1.5, 4.2, 2.9}
   small, big := min_and_max(nums)
   fmt.Println(small, big) // prints 1.5 4.2
   ```

## Question 4

What does `defer` do? What is it commonly useful for in practice? 

What does this code print?

```go
func main() {
	  defer fmt.Println("one")
	  fmt.Println("two")
	  defer fmt.Println("three")
	  fmt.Println("four")
}
```

## Question 5

Write a Go program that uses each of the following functions. In the comments,
give clear descriptions of what each function does:

- `make([]T, len, cap)`
- `append(s, elem)`
- `append(s, elems...)`
- `copy(dst, src)`
- `len(s)`
- `cap(s)`
- `clear(s)` (sets the elements of slice `s` to the zero value; length remains
  unchanged)

## Question 6

Write a function called `countAll(words []string)` that takes a non-empty slice
of strings as input, and returns a map of the counts of all the words in the
slice. Test it on a few different slices of words.

Then write a function called `mostCommon(words []string)` that takes a non-empty
slice, and returns a new `string` slice containing just the strings that occur
the max number of times in `words`. The returned slice should be sorted
alphabetically and have no duplicates.

For example, if `words` is:

```go
{"cat", "bunny", "dog", "cat", "bird", "dog", "bird", "bird", "cat"}
```

Then `mostCommon(words)` should return: 

```go
{"bird", "cat"}
```

That's because `"bird"` occurs 3 times, and `"cat"` occurs 3 times, and no other
word occurs more than 3 times.

## Question 7

Write a function called `compose(f, g)` where:

- `f` and `g` are functions that take a single `int` as input, and return a
  single `int` as output.

- `compose(f, g)` returns a new function that first applies `g` to its input,
  and then applies `f` to the result.

For example, if `f` and `g` are defined as:

```go
func f(x int) int {
  return x + 1
}

func g(x int) int {
  return x * 2
}
```

Then `compose(f, g)` should return a new function that first applies `g` to its
input, and then applies `f` to the result:

```go
h := compose(f, g)
fmt.Println(h(5)) // calculates (5 * 2) + 1, and so prints 11

h = compose(g, f)
fmt.Println(h(5)) // calculates (5 + 1) * 2, and so prints 12
```

## Question 8

Implement a new Go type called `Song` that represents the title and artist of a
song. Then write a `String` method for `Song` that returns a string in the
format "<Title> by <Artist>". 

For example:

```go
song := Song{title: "Bohemian Rhapsody", artist: "Queen"}
fmt.Println(song) // prints: "Bohemian Rhapsody by Queen"
```

## Question 9

Write an interface called `Namer` that has two methods: `GetName()` and
`SetName(name string)`.

Then create these types that implement that `Namer` interface:

- type `Person`, has a `name` field and `age` field
- type `Pet`, which has a `name` field and a `type` field (e.g. "dog", "cat",
  "fish", etc.)

Write a function called `testNamer(n Namer)` that takes a value of type `Namer`
as input, and prints the name of the object, then sets it to a new name, and
prints the new name.

## Question 10

Using Go generics, implement a stack data structure that can store any type `T`,
and implement it using slices (instead of a linked list as in the sample code).
