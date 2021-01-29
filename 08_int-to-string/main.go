package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 75

	y := string(x)

	fmt.Printf("%v - %T\n", x, x)
	fmt.Printf("%v - %T\n", y, y)

	z := strconv.Itoa(x)

	fmt.Printf("%v - %T", z, z)
}
