package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addAnother(x, y int) int {
	return x + y
}

//multiple result
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (y, x int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add(22, 66))
	a, b := "hello", "world"
	a, b = swap(a, b)
	fmt.Println(a, b)
	fmt.Println(split(17))

}
