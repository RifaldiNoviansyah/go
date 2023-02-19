package main

import (
	"fmt"
)

// var can be user outside function
var testString string
var testInt int
var testBool bool
var lastEducation string

func main() {
	// int -  stores integer(whole number), such as 123 or -123
	// float32 - stores float point number, with decimal, such as 19.99 or - -19.99
	// string - stores text, such as "Hello Rifaldi" , string value are surrounded by double quotes
	// bool - stores values with two states : true or false

	// different float16, float32, float64
	// data = 0.123456789121212
	// 16bit: 0.1235 // rounding up get 4 number behind koma
	// 32bit: 0.12345679 // rounding up get 8 number behind koma
	// 64bit: 0.12345678912121212 // rounding up get 17 number behind koma

	var firstName string = "Rifaldi"                // type string
	var lastName = "Noviansyah"                     // type is inferred
	age := 25                                       // type is inferred and := must be inside func
	lastEducation = "Universitas Kompter Indonesia" // variabel declare outside func and separately

	// multiple variabel
	var multipleVariabelSatu, multipleVariabelDua, multipleVariabelTiga, multipleVariabelEmpat int = 1, 2, 3, 4
	multipleVariabelLima, multipleVariabelEnam := 5, "World !"

	var (
		multipleVariabelSeven  int    // 19 digit like int 64
		multipleVariabelEight  int8   = 123
		multipleVariabelNine   int16  = 12345
		multipleVariabelTen    int32  = 1234567891          // 10 digit
		multipleVariabelEleven int64  = 1234567891234567891 //19 digit
		multipleVariabelTwelve string = "Hello"
	)

	// constan variabel
	const PI float32 = 3.14 // namae variabel constan must be uppercase
	// multiple constant variabel
	const (
		multipleVariabelConstanSatu        = 1
		multipleVariabelConstanDua  string = "hi"
	)

	fmt.Println("first name", firstName)
	fmt.Println("last name", lastName)
	fmt.Println("age", age)
	fmt.Println("last education", lastEducation)
	fmt.Println("-----------------")
	fmt.Println("testString is = ", testString)
	fmt.Println("testInt is = ", testInt)
	fmt.Println("testBool is = ", testBool)
	fmt.Println(("---------------"))
	fmt.Println("multipe variabel Satu = ", multipleVariabelSatu)
	fmt.Println("multipe variabel Dua = ", multipleVariabelDua)
	fmt.Println("multipe variabel Tiga = ", multipleVariabelTiga)
	fmt.Println("multipe variabel Empat = ", multipleVariabelEmpat)
	fmt.Println("multipe variabel Lima = ", multipleVariabelLima)
	fmt.Println("multipe variabel Enam = ", multipleVariabelEnam)
	fmt.Println("multipe variabel Seven = ", multipleVariabelSeven)
	fmt.Println("multipe variabel Eight = ", multipleVariabelEight)
	fmt.Println("multipe variabel Nine = ", multipleVariabelNine)
	fmt.Println("multipe variabel Ten = ", multipleVariabelTen)
	fmt.Println("multipe variabel Eleven = ", multipleVariabelEleven)
	fmt.Println("multipe variabel Twelve = ", multipleVariabelTwelve)
	fmt.Println(("---------------"))
	fmt.Println("constan variable", PI)
	fmt.Println("multiple variable constat satu", multipleVariabelConstanSatu)
	fmt.Println("multiple variable constat dua", multipleVariabelConstanDua)
}
