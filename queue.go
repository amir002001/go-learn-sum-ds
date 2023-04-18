package main

import (
	"errors"
	"fmt"
)

func TestQueue() {
	queue := Queue[int]{}
	queue.Init()
	queue.Enqueue(3)
	fmt.Println("enqueued 3")

	queue.Enqueue(4)
	fmt.Println("enqueued 4")

	val, err := queue.Dequeue()
	fmt.Printf("dequeued %d\nlength is %d\n", *val, queue.Length())
	if err != nil {
		panic(err)
	}

	queue.Enqueue(3)
	fmt.Println("enqueued 3")

	queue.Enqueue(5)
	fmt.Println("enqueued 5")

	val, err = queue.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued %d\nlength is %d\n", *val, queue.Length())

	val, err = queue.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued %d\nlength is %d\n", *val, queue.Length())

	val, err = queue.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued %d\nlength is %d\n", *val, queue.Length())

	val, err = queue.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued %d\nlength is %d\n", *val, queue.Length())

	val, err = queue.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued %d\nlength is %d\n", *val, queue.Length())
}

type Node[T any] struct {
	Element *T
	Next    *Node[T]
}

type Queue[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	length int
}

func (queue Queue[T]) Init() {
	queue.length = 0
}

func (queue *Queue[T]) Enqueue(element T) {
	node := Node[T]{&element, nil}
	queue.length++
	if queue.Tail == nil {
		queue.Head = &node
		queue.Tail = &node
		return
	}
	queue.Tail.Next = &node
	queue.Tail = queue.Tail.Next
}

func (queue *Queue[T]) Dequeue() (*T, error) {
	if queue.Head == nil {
		return nil, errors.New("can't dequeue out of empty queue")
	}
	retVal := queue.Head
	queue.Head = queue.Head.Next
	retVal.Next = nil
	queue.length--
	return retVal.Element, nil
}

func (queue *Queue[T]) Peek(node Node[T]) *T {
	return queue.Head.Element
}

func (queue *Queue[T]) Length() int {
	return queue.length
}
