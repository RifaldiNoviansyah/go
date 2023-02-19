package main

import (
	"fmt"
)

func main() {
	// \n = for new line
	// %s = for string
	// %d = for intger
	// %v = for float
	// %v = for value for everyting
	// %T = for know the type

	// variabel
	firstName := "Rifaldi"
	lastName := "Noviansyah"
	age := 25
	length := 2.5
	// "Print" This can't format anything, it simply takes a string and print it
	fmt.Print("Hello")
	fmt.Print("Rifaldi")
	fmt.Print("Hello", firstName)
	// fmt.Print("Hellor %s", firstName) // error masukan %s %d etc hanya bisa di printf
	//  "Print Line" same thing as Print() however it will append a newline character \n at the end.
	fmt.Println("Hello Rifaldi")
	fmt.Println("Hello", firstName)
	// fmt.Println("Hello %s", firstName) // error masukan %s %d etc hanya bisa di printf
	// "Print Formatter" this function allows you to format numbers, variables and strings into the first string parameter you give it
	fmt.Printf("1. My Name is: %s %s %d %v \n", firstName, lastName, age, length)
	fmt.Printf("2. My Name is: %v %v %v %v \n", firstName, lastName, age, length)
	fmt.Printf("2. My Name is: %T %T %T %T \n", firstName, lastName, age, length)
}
