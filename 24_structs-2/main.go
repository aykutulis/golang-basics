package main

import "fmt"

type person struct {
	height    float64
	weight    float64
	hairColor string
}

type employee struct {
	person
	profession string
	salary     int
}

func main() {

	/* e1 := employee{
		person: person{
			height:    175.5,
			weight:    85.3,
			hairColor: "yellow",
		},
		profession: "Mining engineer",
		salary:     5000,
	}

	fmt.Println(e1)

	e2 := employee{
		person: person{
			height:    165.5,
			hairColor: "yellow",
		},
		profession: "Mining engineer",
	}

	fmt.Println(e2) // zero value */

	/* p1 := person{
		height: 170.7,
	}

	fmt.Println(p1) */

	// Anonymous Struct
	product := struct {
		name  string
		price int
	}{name: "Computer", price: 5000}

	fmt.Println(product)

}
