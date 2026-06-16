// generic_stack/main.go

package main

import "fmt"

//
// Lowercase names, like node, are package-private, i.e. only visible within the
// package. Uppercase names, like Stack, are public, i.e. visible outside the
// package.
//

// "any" means that T can be any type: there are no constraints on T
type node[T any] struct {
	data T
	next *node[T] // next is pointer to a node of type T
}

// Stack stores items of type T as a linked list. head points to the top of the
// stack.
type Stack[T any] struct {
	head *node[T]
}

// returns a new empty stack
func MakeStack[T any]() Stack[T] {
	return Stack[T]{head: nil}
}

// returns true if the stack is empty
func (s Stack[T]) IsEmpty() bool {
	return s.head == nil
}

// adds a new item to the top of the stack
func (s *Stack[T]) Push(val T) {
	s.head = &node[T]{data: val, next: s.head}
}

// removes the top item from the stack and returns it; assumes the stack is not
// empty
func (s *Stack[T]) Pop() T {
	val := s.head.data
	s.head = s.head.next
	return val
}

// removes the top item from the stack and returns it and true; returns the zero
// value of type T and false if the stack is empty
func (s *Stack[T]) SafePop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	val := s.head.data
	s.head = s.head.next
	return val, true
}

// Go has a garbage collector, so we don't need to manually free memory by
// calling Pop.
func (s *Stack[T]) Clear() {
	s.head = nil
}

// returns the top item from the stack without removing it; assumes the stack is
// not empty
func (s Stack[T]) Peek() T {
	return s.head.data
}

// returns the top item from the stack without removing it; returns the zero
// value of type T and false if the stack is empty
func (s Stack[T]) SafePeek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.head.data, true
}

// prints the stack to the console
func (s Stack[T]) Print() {
	if s.IsEmpty() {
		fmt.Println("<empty stack>")
	} else {
		fmt.Print("Stack: top=")
		for p := s.head; p != nil; p = p.next {
			fmt.Print(p.data, " ")
		}
		fmt.Println()
	}
}

func demoInt() {
	s := MakeStack[int]()
	s.Print()
	fmt.Println(s.IsEmpty())
	s.Push(1)
	s.Print()
	s.Push(2)
	s.Print()
	s.Push(3)
	s.Print()

	fmt.Println("popping: ", s.Pop())
	s.Print()
	fmt.Println("popping: ", s.Pop())
	s.Print()
	fmt.Println("popping: ", s.Pop())
	s.Print()
}

func demoString() {
	s := MakeStack[string]()
	s.Print()
	fmt.Println(s.IsEmpty())
	s.Push("hello")
	s.Print()
	s.Push("world")
	s.Print()
	s.Push("!")
	s.Print()

	fmt.Println("popping: ", s.Pop())
	s.Print()
	fmt.Println("popping: ", s.Pop())
}

func main() {
	// demoInt()
	demoString()
}
