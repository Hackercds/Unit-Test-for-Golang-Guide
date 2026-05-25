package ds_test

import (
	"DeepTest/ds"
	"fmt"
)

func ExampleNewStack() {
	s := ds.NewStack()
	fmt.Println(s.IsEmpty())
	// Output: true
}

func ExampleStack_Push() {
	s := ds.NewStack()
	s.Push(10)
	s.Push(20)
	val, _ := s.Peek()
	fmt.Println(val)
	// Output: 20
}

func ExampleStack_Pop() {
	s := ds.NewStack()
	s.Push(10)
	s.Push(20)
	val, _ := s.Pop()
	fmt.Println(val)
	// Output: 20
}

func ExampleNewQueue() {
	q := ds.NewQueue()
	fmt.Println(q.IsEmpty())
	// Output: true
}

func ExampleQueue_Enqueue() {
	q := ds.NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	val, _ := q.Peek()
	fmt.Println(val)
	// Output: 10
}

func ExampleQueue_Dequeue() {
	q := ds.NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	val, _ := q.Dequeue()
	fmt.Println(val)
	// Output: 10
}
