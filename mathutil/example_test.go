package mathutil_test

import (
	"DeepTest/mathutil"
	"fmt"
)

func ExampleFibonacci() {
	fmt.Println(mathutil.Fibonacci(10))
	// Output: 55
}

func ExampleFibonacci_negative() {
	fmt.Println(mathutil.Fibonacci(-5))
	// Output: -1
}

func ExampleFactorial() {
	fmt.Println(mathutil.Factorial(5))
	// Output: 120
}

func ExampleFactorial_zero() {
	fmt.Println(mathutil.Factorial(0))
	// Output: 1
}

func ExampleGCD() {
	fmt.Println(mathutil.GCD(48, 18))
	// Output: 6
}

func ExampleIsPrime() {
	fmt.Println(mathutil.IsPrime(17))
	fmt.Println(mathutil.IsPrime(100))
	// Output:
	// true
	// false
}
