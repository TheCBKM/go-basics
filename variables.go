package main

import "fmt"

var c, python, java bool
var str string
var a1, a2, a3 = true, "wow", 123

func main() {
	short := "short declaration"
	var i int
	var k, l = 1, 2
	fmt.Println(i, str, c, python, java)
	fmt.Println(k, l)
	fmt.Println(a1, a2, a3)
	fmt.Println(short)
}
