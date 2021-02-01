package main

import "fmt"

func main() {

	x := 5
	y := &x

	fmt.Println(x, y)
	fmt.Println(x, *y)
	fmt.Printf("%T - %v\n", y, y)

	*y = 33

	fmt.Println(x)

}
