package ds

import "errors"

// Queue is a FIFO data structure backed by a slice.
type Queue struct {
	data  []int
	front int
}

// NewQueue creates a new empty Queue.
func NewQueue() *Queue {
	return &Queue{data: make([]int, 0), front: 0}
}

// Enqueue adds a value to the back of the queue.
func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

// Dequeue removes and returns the value at the front of the queue.
// Returns an error if the queue is empty.
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("dequeue from empty queue")
	}
	val := q.data[q.front]
	q.front++
	if q.front > len(q.data)/2 {
		q.data = q.data[q.front:]
		q.front = 0
	}
	return val, nil
}

// Peek returns the value at the front of the queue without removing it.
// Returns an error if the queue is empty.
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("peek from empty queue")
	}
	return q.data[q.front], nil
}

// IsEmpty returns true if the queue contains no elements.
func (q *Queue) IsEmpty() bool {
	return q.front >= len(q.data)
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.data) - q.front
}
