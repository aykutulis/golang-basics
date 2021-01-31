package main

import "fmt"

func main() {

	/* 	city1 := "istanbul"
	   	city2 := "londra"
	   	city3 := "new york"
	   	city4 := "berlin"
	   	fmt.Println(city1, city2, city3, city4) */

	//cities := [4]string{"istanbul", "londra", "new york", "berlin"}

	/* 	cities := []string{"istanbul", "londra", "new york", "berlin"}
	   	fmt.Println(cities)
	   	fmt.Println(cities[0]) // zero based index
	   	fmt.Println(cities[3])
	   	fmt.Println(len(cities)) */

	/* 	var myArr [5]int
	   	fmt.Println(myArr)
	   	//myArr[0] = "istanbul"
	   	myArr[0] = 100
	   	fmt.Println(myArr)
	   	myArr[len(myArr)-1] = 200
	   	fmt.Println(myArr) */

	/* 	var myArr [4]int
	   	var myArr2 [5]int */

	/* 	if myArr == myArr2 {
		fmt.Println("Hello!")
	} */

	/* 	fmt.Println(myArr)
	   	fmt.Println(myArr2) */

	/* 	cities := [4]string{"istanbul", "londra", "new york", "berlin"}
	   	for i := 0; i < len(cities); i++ {
	   		fmt.Println(i, cities[i])
	   	}
	   	cities[0] = "BOLU"
	   	fmt.Println()
	   	for i := 0; i < len(cities); i++ {
	   		fmt.Println(i, cities[i])
	   	} */

	/* 	myArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	   	myArr = mySquare(myArr) // First Class Functions
	   	fmt.Println(myArr) */

	// FOR --- RANGE

	cities := [4]string{"istanbul", "londra", "new york", "berlin"}

	/* 	for index, city := range cities {
		fmt.Println(index, city)
	} */

	for _, city := range cities {
		fmt.Println(city)
	}
}
