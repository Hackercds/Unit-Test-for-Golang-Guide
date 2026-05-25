package mathutil

// Fibonacci returns the nth Fibonacci number (0-indexed: F(0)=0, F(1)=1).
// Returns -1 for negative input.
func Fibonacci(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
