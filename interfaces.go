package main

import "fmt"

type I interface {
	M()
}

type shape interface {
	area() float64
}

type T struct {
	S string
}

type circle struct {
	radius float64
}

type rect struct {
	length  float64
	breadth float64
}

func (s circle) area() float64 {
	return 3.14 * (s.radius * s.radius)
}

func (s rect) area() float64 {
	return s.length * s.breadth
}

//M This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()

	rec := rect{3, 5}
	cir := circle{4}

	fmt.Println(rec.area())
	fmt.Println(cir.area())

	shapes := []shape{rec, cir}
	fmt.Println(shapes[0].area())
	fmt.Println(shapes[1].area())
}
