package ds_test

import (
	"DeepTest/ds"
	"testing"
)

func BenchmarkStackPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ds.NewStack()
		for j := 0; j < 1000; j++ {
			s.Push(j)
		}
	}
}

func BenchmarkStackPop(b *testing.B) {
	const setupSize = 10000
	s := ds.NewStack()
	for j := 0; j < setupSize; j++ {
		s.Push(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if s.IsEmpty() {
			for j := 0; j < setupSize; j++ {
				s.Push(j)
			}
		}
		s.Pop()
	}
}

func BenchmarkQueueEnqueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := ds.NewQueue()
		for j := 0; j < 1000; j++ {
			q.Enqueue(j)
		}
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	const setupSize = 10000
	q := ds.NewQueue()
	for j := 0; j < setupSize; j++ {
		q.Enqueue(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if q.IsEmpty() {
			for j := 0; j < setupSize; j++ {
				q.Enqueue(j)
			}
		}
		q.Dequeue()
	}
}
