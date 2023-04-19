package main

import (
	"errors"
	"fmt"
)

func TestStack() {
	stack := Stack[int]{}
	stack.Push(1)
	fmt.Printf("pushed 1\n")
	stack.Push(8)
	fmt.Printf("pushed 8\n")
	stack.Push(9)
	fmt.Printf("pushed 9\n")

	val, err := stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("popped %d. length is %d\n", *val, stack.Length())

	val, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("popped %d. length is %d\n", *val, stack.Length())

	stack.Push(3)
	fmt.Printf("pushed 3\n")

	val, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("popped %d. length is %d\n", *val, stack.Length())

	val, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("popped %d. length is %d/n", *val, stack.Length())

	val, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("popped %d. length is %d\n", *val, stack.Length())

	val, err = stack.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("popped %d. length is %d\n", *val, stack.Length())
}

type Stack[T any] struct {
	Head   *Node[T]
	length int
}

func (stack *Stack[T]) Push(value T) {
	node := Node[T]{&value, nil}
	if stack.length == 0 {
		stack.Head = &node
	} else {
		node.Next = stack.Head
		stack.Head = &node
	}
	stack.length++
}

func (stack *Stack[T]) Pop() (*T, error) {
	if stack.length == 0 {
		return nil, errors.New("can't pop out of empty stack")
	}
	stack.length--
	tmp := stack.Head
	fmt.Printf("inside pop %v\n", *stack.Head.Element)
	stack.Head = stack.Head.Next
	tmp.Next = nil
	return tmp.Element, nil
}

func (stack *Stack[T]) Peek() *T {
	return stack.Head.Element
}

func (stack *Stack[T]) Length() int {
	return stack.length
}
