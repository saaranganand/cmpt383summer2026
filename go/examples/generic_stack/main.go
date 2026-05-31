// generic_stack/main.go

package main

import "fmt"

//
// Lowercase names, like node, are package-private, i.e. only visible within the
// package. Uppercase names, like Stack, are public, i.e. visible outside the
// package.
//

// T any means that T can be any type: there are no constraints on T.
type node[T any] struct {
	val  T
	next *node[T]
}

// Stack stores items of type T as a linked list.
type Stack[T any] struct {
	head *node[T]
}

func MakeEmpty[T any]() Stack[T] {
	// by not explicitly initializing val here, val gets assigned the zero value
	// of type T, which is what we want
	return Stack[T]{head: nil}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

// adds a new item to the top of the stack
func (s *Stack[T]) Push(val T) {
	s.head = &node[T]{val: val, next: s.head}
}

// removes the top item from the stack and returns it; assumes the stack is not
// empty
func (s *Stack[T]) Pop() T {
	val := s.head.val
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
	val := s.head.val
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
func (s *Stack[T]) Peek() T {
	return s.head.val
}

// returns the top item from the stack without removing it; returns the zero
// value of type T and false if the stack is empty
func (s *Stack[T]) SafePeek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.head.val, true
}

// prints the stack to the console
func (s *Stack[T]) Print() {
	if s.IsEmpty() {
		fmt.Println("<empty stack>")
	} else {
		fmt.Print("Stack: top=")
		for p := s.head; p != nil; p = p.next {
			fmt.Print(p.val, " ")
		}
		fmt.Println()
	}
}

func main() {
	s1 := MakeEmpty[int]()
	s1.Print()
	fmt.Println(s1.IsEmpty())
	s1.Push(1)
	s1.Print()
	s1.Push(2)
	s1.Print()
	s1.Push(3)
	s1.Print()

	fmt.Println("popping: ", s1.Pop())
	s1.Print()
	fmt.Println("popping: ", s1.Pop())
	s1.Print()
	fmt.Println("popping: ", s1.Pop())
	s1.Print()
}
