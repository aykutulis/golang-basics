/* package main
import (
	"errors"
	"fmt"
	"math"
)
func main() {
	result, err := sRoot(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("You should have entered possitive number")
	}
	return math.Sqrt(num), nil
} */

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız", file)
	}

}
