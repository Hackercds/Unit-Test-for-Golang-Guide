package mathutil_test

import (
	"DeepTest/mathutil"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	b.Run("n=10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.Fibonacci(10)
		}
	})
	b.Run("n=20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.Fibonacci(20)
		}
	})
	b.Run("n=50", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.Fibonacci(50)
		}
	})
}

func BenchmarkFactorial(b *testing.B) {
	b.Run("n=10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.Factorial(10)
		}
	})
	b.Run("n=20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.Factorial(20)
		}
	})
}

func BenchmarkGCD(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.GCD(48, 18)
		}
	})
	b.Run("large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.GCD(123456, 789012)
		}
	})
}

func BenchmarkIsPrime(b *testing.B) {
	b.Run("small_prime", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.IsPrime(17)
		}
	})
	b.Run("large_prime", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.IsPrime(9973)
		}
	})
	b.Run("composite", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mathutil.IsPrime(9999)
		}
	})
}
