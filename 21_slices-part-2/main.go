package main

import "fmt"

func main() {
	/* underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(underArray)

	mySlc := underArray[:5]

	mySlc[0] = 100

	fmt.Println(mySlc)
	fmt.Println(underArray) */

	/* mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)

	mySlc = append(mySlc, 4, 5)
	fmt.Println(mySlc)
	fmt.Printf("%#v", mySlc)

	mySlc2 := append(mySlc, 4, 5)
	fmt.Println(mySlc2)

	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2) */

	/* mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5}

	mySlc = append(mySlc, mySlc2...)
	fmt.Println(mySlc) */

	/* mySlc := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mySlc)
	mySlc = mySlc[2 : len(mySlc)-3]
	fmt.Println(mySlc) */

	var mySlc3 []int
	fmt.Printf("%#v", mySlc3) // nil

	fmt.Println()

	mySlc4 := make([]int, 0)
	fmt.Printf("%#v", mySlc4) // {}

}
