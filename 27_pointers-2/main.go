/* package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x)

	double(x)
	fmt.Println(x)

}

func double(num int) {
	num *= 2
	fmt.Println(num)
} */

package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x)

	double(&x)

	fmt.Println(x)

}

func double(num *int) {
	*num *= 2
	fmt.Println(num)
	fmt.Println(*num)
}
