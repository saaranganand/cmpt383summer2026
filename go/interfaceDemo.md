# Go's Stringer Interface

In Go, an **interface** type is a collection of **method signatures**. For
example, the standard interface `Stringer` is defined like this:

```go
type Stringer interface {
    String() string // method signature
    // the methods name is String, it takes no arguments, and returns a string
}
```

`Stringer` happens to have only one method, but in general interfaces can have
any number of methods.

A type is said to **implement an interface** if it has methods for all the
methods listed in it. So a type implements `Stringer` if it implements a method
that looks like this:

```go
func (this T) String() string {
	// ... returns a string representation of of T ...
}
```

The parameter `(this T)` that appears before the name is called the **receiver
argument** of the method. Importantly, the receiver of a method does *not*
appear in an interface. This allows different types to implement their own
version of `String()`.

In this example the receiver argument is named `this`, as in C++. But in Go, the
receiver variable can be named by the programmer. It's usually good to name it
something descriptive related to the type of the receiver, e.g. for a
`BankAccount` type we might call the receiver `ba`.

## Implementing Your Own Stringer

Suppose you implement your own type `Person`:

```go
type Person struct {
    name string
    age int
}

func main() {
    p := Person{"Bob", 65}
    fmt.Println(p)  // "{Bob 65}"
}
```

When `p` is printed with a standard Go function like `fmt.Println`, it gets
printed in a default format: `"{Bob 65}"`.

What if you want `Person` objects printed in some other format, say `Person(Bob,
65)`? You can do that by adding a `String()` method like this:

```go
type Person struct {
    name string
    age int
}

func (p Person) String() string {
    return fmt.Sprintf("Person(%v, %v)", p.name, p.age)
    // fmt.Sprintf works like fmt.Printf but returns a string
}

func main() {
    p := Person{"Bob", 65}
    fmt.Println(p)  // "Person(Bob, 65)"
}
```

Importantly, the `String()` method must *exactly* match the signature in
`Stringer`. If it does, then the type `Person` is said to **implement** the
`Stringer` interface. Now when `fmt.Println(p)` is called `"Person(Bob, 65)"` is
printed.

Suppose you have different ways that you want to print a `Person` object. The
`String()` method is generally used as the default way to convert
 value to a string. For other ways of printing it you should use other methods, e.g.:

```go
func (p Person) JSON() string {
	return fmt.Sprintf("{\"name\": \"%v\", \"age\": %v}", p.name, p.age)
}

// fmt.Println(p)        // prints: "Person(Bob, 65)"
// fmt.Println(p.JSON()) // prints: "{\"name\": \"Bob\", \"age\": 65}"
```

## Counter: An Interface for Counting

Lets create our own interface. `Counter` is an interface that represents an
integer value that can incremented, read, and reset:

```go
type Counter interface {
    incr(n int)     // increment the counter by n
    getCount() int  // the current value of the counter
    reset()         // set the counter to 0
}
```

To implement `Counter`, we need to create a new type that has it three methods.
For instance:

```go
type NamedCount struct {
	name  string
	count int
}

// increment the counter by n
// nc is passed as a pointer since the underlying count is modified
func (nc *NamedCount) incr(n int) {
	nc.count += n
}

// the current value of the counter
func (nc NamedCount) getCount() int {
	return nc.count
}

// set the counter to 0
// nc is passed as a pointer since the underlying count is modified
func (nc *NamedCount) reset() {
	nc.count = 0
}

// All we know about the passed-in value c is that it implements the Counter
// interface, and so the incr, getCount, and reset methods can be called on
// it.
func testCounter(c Counter) {
    fmt.Println(c.getCount()) // 0
    c.incr(1)
    c.incr(1)
    fmt.Println(c.getCount()) // 2
    c.reset()
    fmt.Println(c.getCount()) // 0

    fmt.Println(c)
}

func main() {
	a := &NamedCount{"Test counter", 0}
    testCounter(a)
}
```

Notice that the type of parameter `c` in `testCounter` is `Counter`, i.e. it is
an interface. Thus `testCounter` knows `c` has the three methods listed in
`Counter`.

An interesting detail is that `NamedCount` *implicitly* implements the
`Stringer` interface. Nowhere do we explicitly say in the program that
`NamedCount` implements `Counter`. Go automatically infers it because it has the
three methods listed in `Counter`. 


### The `String()` Method

Look at the last line of `testCounter`: `fmt.Println(c)`. This calls the
`String()` method of `c`. But where does this `String()` method come from?
`Counter` doesn't have a `String()` method. 

The basic answer is that the *underlying* type of `c` is `NamedCount`, and
`NamedCount` implements the `Stringer` interface (by default). Go calls the
`String()` method of the underlying type when `fmt.Println(c)` is called.

If we wanted to, we could add a `String()` method to `NamedCount`:

```go
// implements the Stringer interface 
func (nc NamedCount) String() string {
	return fmt.Sprintf("counter%v: %v", nc.name, nc.count)
}
```

Go automatically infers that `NamedCount` implements the `Stringer` interface
because it has this `String()` method. When `fmt.Println(c)` is called, this
method is called, and so `"counter<name>: <count>"` is printed.

### Re-using `testCounter`

If we create another type that implements `Counter`, we can re-use
`testCounter`, e.g.:

```go
// ... previous code ...

type UndoableCounter struct {
	count int
	prevCount int
}

func (uc *UndoableCounter) incr(n int) {
	uc.count += n
	uc.prevCount = uc.count
}

func (uc *UndoableCounter) getCount() int {
	return uc.count
}

func (uc *UndoableCounter) reset() {
	uc.count = uc.prevCount
}

// not part of the Counter interface
func (uc *UndoableCounter) undo() {
	uc.count = uc.prevCount
}

// new main function
func main() {
	fmt.Println("Testing NamedCounter...")
	a := &NamedCount{"Test counter", 0}
    testCounter(a)

	fmt.Println("\nTesting UndoableCounter...")
	b := &UndoableCounter{0, 0}
    testCounter(b)
}
```

Notice that `UndoableCounter` has the method `undo` that is not part of the
`Counter` interface. `undo` *cannot* be called in `testCounter` because `c` is
of type `Counter`, and `Counter` does not have a `undo` method. For instance, if
the underlying type of `c` was `NamedCount`, then there is no `undo` method to
call.

### Non-struct Types

Non-struct types can also implement an interface. For example:

```go
type BasicCount int

// increment the counter by n
func (bc *BasicCount) incr(n int) { 
    *bc += BasicCount(n)
}

// the current value of the counter
func (bc BasicCount) getCount() int { 
    return int(bc)
}

// set the counter to 0
func (bc *BasicCount) reset() {
    *bc = 0
}

// ...

func main() {
    b := BasicCount(0)
    testCounter(&b) // b is passed as pointer because Counter has pointer
}                   // receivers
```

`BasicCount` has the same internal implementation as an `int`, but it's *not*
the same type as an `int`, i.e. `int` and `BasicCount` *cannot* be used
interchangeably. Instead, you must explicitly convert between the types, as
shown in `BasicCount`'s `getCount` method.

Note that `incr` and `reset` both use a *pointer* receiver. That's so the
value can be changed in-place. If you don't use a pointer, then a copy is made
of the receiver parameter.

See [interface_Counter](examples/interface_Counter/main.go) for code used in
this note.
