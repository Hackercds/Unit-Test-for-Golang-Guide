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
	s := ds.NewStack()
	for j := 0; j < b.N+10000; j++ {
		s.Push(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
	q := ds.NewQueue()
	for j := 0; j < b.N+10000; j++ {
		q.Enqueue(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
