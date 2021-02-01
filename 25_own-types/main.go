package main

import "fmt"

type kilometer float64
type mile float64

func main() {

	k1 := kilometer(5)

	m1 := mile(10)

	fmt.Println(toMile(k1))
	fmt.Println(toKilometer(m1))

}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

func toMile(k kilometer) mile {
	return mile(k * 0.62)
}
