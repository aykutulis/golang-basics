package main

import "fmt"

func main() {

	grade := -3
	switch grade {
	case 5:
		fmt.Println("Excellent")
	case 4:
		fmt.Println("Good")
	case 3:
		fmt.Println("Avarage")
	case 2:
		fmt.Println("Not bad")

	case 1:
		fmt.Println("Bad")

	default:
		fmt.Println("Invalid")
	}
}
