package main

import "fmt"

func testDeffer() string {
	defer fmt.Println("world")

	fmt.Println("hello")
	return "my"
}

// Deferred function calls are pushed onto a stack.
// When a function returns,
// its deferred calls are executed in last-in-first-out order.

func withStck() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	fmt.Println(testDeffer())
	withStck()
}
