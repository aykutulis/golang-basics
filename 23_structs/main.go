package main

import "fmt"

func main() {

	type employee struct {
		name      string
		age       int
		isMarried bool
		tech      []string
	}

	e1 := employee{
		name:      "Aykut",
		age:       29,
		isMarried: false,
		tech: []string{
			"JS",
			"Golang",
			"Node",
			"React",
			"Nestjs",
			"Nextjs",
			"MongoDB",
			"PostgreSQL",
		},
	}

	e1.age = 30

	fmt.Println(e1)
	fmt.Println(e1.age)
	fmt.Println(e1.tech)
	fmt.Println(e1.tech[3])

}
