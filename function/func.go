package main

import (
	"fmt"
)

func myName() {
	fmt.Println("My Name : Rifaldi Noviansyah")
}

func myAge(age int) {
	fmt.Println("My Age : ", age)
}

func myIpk(x float32, y float32) (result float32) {
	return x + y
}

func main() {
	myName()
	myAge(25)
	fmt.Println(myIpk(1.0, 2.25))
}
