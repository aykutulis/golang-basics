package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
}

func area(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}

func main() {

	c1 := circle{5}
	chan1 := make(chan float64)

	go area(c1, chan1)

	fmt.Printf("%2f\n", <-chan1)

	c1.display()

}
