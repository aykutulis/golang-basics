package main

import "fmt"

func main() {

	/* 	x := 10
	   	if x := -5; x < 0 {
	   		fmt.Println(x, "negative")
	   	} else if x%2 == 0 {
	   		fmt.Println(x, "even")
	   	} else {
	   		fmt.Println(x, "odd")
	   	}
		   fmt.Println(x) */

	if x := -25; x < 0 {
		fmt.Println(x, "negative")
		fmt.Println("Hello!")
	} else {
		if x%2 == 0 {
			fmt.Println(x, "even")
		} else {
			fmt.Println(x, "odd")
		}
	}

}
