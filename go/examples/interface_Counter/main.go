// interface_Counter/main.go

package main

import "fmt"

//
// The Counter interface represents and integer object that can be incremented,
// read, and reset. 
//
// It does not care about the internal implementation of how this is done.
//
// Any type that has, at least, the three methods listed in the Counter
// interface, is said to implement the Counter interface. Go automatically
// infers this: you don't need to explicitly say that the type implements the
// Counter interface.
//

type Counter interface {
	incr(n int)    // increment the counter by n
	getCount() int // the current value of the counter
	reset()        // set the counter to 0
}

/////////////////////////////////////////////

//
// a counter with a name
//
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

// returns the name of the counter (not part of the Counter interface)
func (nc NamedCount) getName() string {
	return nc.name
}

// implement the Stringer interface
func (nc NamedCount) String() string {
	return fmt.Sprintf("counter%v: %v", nc.name, nc.count)
}

//
// All we know about the passed-in value c is that it implements the Counter
// interface, and so the incr, getCount, and reset methods can be called on it.
//
// Notice we *cannot* call the getName method on c because it is not part of the
// Counter interface.
//
func testCounter(c Counter) {
	fmt.Println(c.getCount()) // 0
	c.incr(1)
	c.incr(1)
	fmt.Println(c.getCount()) // 2
	c.reset()
	fmt.Println(c.getCount()) // 0

	fmt.Println(c)
}

/////////////////////////////////////////////

//
// a counter that can be undone one level (i.e. can be reset to the previous
// value)
//
type UndoableCounter struct {
	count     int
	prevCount int
}

func (uc *UndoableCounter) incr(n int) {
	uc.prevCount = uc.count
	uc.count += n
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

/////////////////////////////////////////////

//
// This makes a new type that is different than int, although it has the same
// internal implementation as an int.
//
// This is different than typedef in C/C++. typedef does not create a new type,
// but rather creates an alias for a type. The alias is just another name for
// the type, and the original type and its alias can be used interchangeably.
//

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

func main() {
	fmt.Println("Testing NamedCounter...")
	a := &NamedCount{"Test counter", 0}
	testCounter(a)

	fmt.Println("\nTesting UndoableCounter...")
	b := &UndoableCounter{0, 0}
	testCounter(b)

	fmt.Println("\nTesting BasicCount...")
	c := BasicCount(0)
	testCounter(&c)
	// c is passed as a pointer because Counter has pointer receivers
}
