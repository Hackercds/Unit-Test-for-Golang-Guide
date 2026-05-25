package mathutil

// GCD returns the greatest common divisor of a and b using the Euclidean algorithm.
// GCD(0, 0) returns 0 by convention.
func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
