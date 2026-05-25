package ds_test

import (
	"DeepTest/ds"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := ds.NewQueue()
	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}
	if q.Size() != 0 {
		t.Errorf("new queue size should be 0, got %d", q.Size())
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := ds.NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)

	if q.Size() != 2 {
		t.Errorf("size = %d, want 2", q.Size())
	}
	val, err := q.Peek()
	if err != nil {
		t.Fatalf("Peek() error: %v", err)
	}
	if val != 10 {
		t.Errorf("Peek() = %d, want 10", val)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := ds.NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)

	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error: %v", err)
	}
	if val != 10 {
		t.Errorf("Dequeue() = %d, want 10", val)
	}
	if q.Size() != 1 {
		t.Errorf("size = %d, want 1", q.Size())
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error: %v", err)
	}
	if val != 20 {
		t.Errorf("Dequeue() = %d, want 20", val)
	}
	if !q.IsEmpty() {
		t.Error("queue should be empty after dequeuing all elements")
	}
}

func TestQueueDequeueEmpty(t *testing.T) {
	q := ds.NewQueue()
	_, err := q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing empty queue")
	}
}

func TestQueuePeekEmpty(t *testing.T) {
	q := ds.NewQueue()
	_, err := q.Peek()
	if err == nil {
		t.Error("expected error when peeking empty queue")
	}
}

func TestQueueFIFO(t *testing.T) {
	q := ds.NewQueue()
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		q.Enqueue(v)
	}
	for i := 0; i < len(values); i++ {
		val, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue() error: %v", err)
		}
		if val != values[i] {
			t.Errorf("Dequeue() = %d, want %d", val, values[i])
		}
	}
}
