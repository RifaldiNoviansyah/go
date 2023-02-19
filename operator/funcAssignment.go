package main

import (
	"fmt"
)

var tempOne int
var tempTwo int = 5
var tempThree int = 5
var tempFour int = 5
var tempFive = 5

func main() {
	tempOne = 5
	tempTwo += 5
	tempThree -= 5
	tempFour *= 5
	tempFive |= 5 // user binary representation

	fmt.Printf("tempOne = 5 = %d\n", tempOne)
	fmt.Printf("tempTwo += 5 = %d\n", tempTwo)
	fmt.Printf("tempThree -= 5 = %d\n", tempThree)
	fmt.Printf("tempFour *= 5 = %d\n", tempFour)
	fmt.Printf("tempFive |= 5 = %d\n", tempFive)
}
