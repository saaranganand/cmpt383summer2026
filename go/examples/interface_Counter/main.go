// interface_Counter/main.go

package main

import "fmt"

type Counter interface {
	incr(n int)    // increment the counter by n
	getCount() int // the current value of the counter
	reset()        // set the counter to 0
}

/////////////////////////////////////////////

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

// implement the Stringer interface
func (nc NamedCount) String() string {
	return fmt.Sprintf("counter%v: %v", nc.name, nc.count)
}

// All we about the passed-in value c is that it implements the Counter
// interface, and so the incr, getCount, and reset methods can be called on it.
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

type UndoableCounter struct {
	count     int
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

/////////////////////////////////////////////

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
