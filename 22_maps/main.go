package main

import "fmt"

func main() {

	/* myMap := map[string]int{
		"Alan":    35,
		"Nikola":  40,
		"Sigmund": 51,
		"Arthur":  65,
	} */

	/* fmt.Println(myMap)
	fmt.Println(myMap["Nikola"])
	fmt.Println(myMap["Haydar"])

	_, ok := myMap["Haydar"]
	fmt.Println(ok)

	myMap["Carl"] = 45
	fmt.Println(myMap)
	fmt.Println(myMap["Carl"]) */

	/* var myMap map[string]int

	fmt.Printf("%#v", myMap) // nil */

	/* myMap := map[string]bool{
		"Alan":   false,
		"Arthur": true,
		"Breur":  true,
	}

	myMap2 := myMap

	delete(myMap2, "Alan")
	fmt.Println(myMap2)
	fmt.Println(myMap) // pass by reference */

	/* myMap := make(map[string]int)
	fmt.Println(myMap)
	fmt.Println(myMap["Test"])

	_, ok := myMap["Test"]
	fmt.Println(ok) */

	myMap := map[string]int{
		"Alan":    35,
		"Nikola":  40,
		"Sigmund": 51,
		"Arthur":  65,
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
