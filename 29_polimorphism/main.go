package main

import "fmt"

type shape interface {
	area() float64
}

func getAreas(shapes ...shape) {

	for _, shape := range shapes {
		fmt.Println("Area:", shape.area())
	}

}

type triangle struct {
	a, h float64
}

func (t triangle) area() float64 {
	return t.a * t.h / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func main() {
	t := triangle{7, 8}
	s := square{5}
	r := rectangle{3, 4}

	getAreas(t, s, r)
}
