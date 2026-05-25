package tttys_test

import (
	"DeepTest/tttys"
	"fmt"
)

func ExampleGreeting() {
	fmt.Println(tttys.Greeting())
	// Output: Hello from tttys!
}

func ExampleInitialize() {
	fmt.Println(tttys.Initialize())
	// Output: tttys initialized
}
