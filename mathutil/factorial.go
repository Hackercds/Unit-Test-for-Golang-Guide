package mathutil

// Factorial returns n! (n factorial = 1*2*...*n).
// Returns 1 for n=0. Returns -1 for negative input.
func Factorial(n int) int {
	if n < 0 {
		return -1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
