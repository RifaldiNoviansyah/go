package main

import (
	"fmt"
)

var day int

func main() {
	day = 5

	switch day {
	case 1, 3, 5:
		fmt.Println("odd day")
	case 2, 4:
		fmt.Println("even day")
	default:
		fmt.Println("invalid day of day number")
	}
}
