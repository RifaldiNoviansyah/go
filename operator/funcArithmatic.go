package main

import (
	"fmt"
)

var addition int8
var subtraction int8
var multiplication int16
var division float32
var increment int = 1
var decrement int = 2
var modulus int

func main() {
	addition = 15 + 12
	subtraction = 15 - 12
	multiplication = 15 * 12
	division = 15.0 / 12.0 // if we want the output there is a number after the comma we must use .0 if no the output will be fulfilled
	increment++
	decrement--
	modulus = 2 % 5

	fmt.Printf("addition = %d\n", addition)
	fmt.Printf("subtraction = %d\n", subtraction)
	fmt.Printf("multiplication = %d\n", multiplication)
	fmt.Printf("division = %v\n", division)
	fmt.Printf("increment = %d\n", increment)
	fmt.Printf("decrement = %d\n", decrement)
	fmt.Printf("modulus = %d\n", modulus)
}
