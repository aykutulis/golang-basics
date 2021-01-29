package main

import "fmt"

/* var packVar = "Package Scope"

var funcVar = "Func(Package) Scope" */

func main() {

	/* 	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	} */

	var name = "Sigmund"
	name, surname := "Anna", "Freud"

	fmt.Println(name, surname)

	/* 	fmt.Println(packVar)
	   	myFunc() */

}

/* func myFunc() {
	fmt.Println(funcVar)
}
*/
