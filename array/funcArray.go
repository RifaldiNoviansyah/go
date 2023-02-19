package main

import (
	"fmt"
)

func main() {

	// array with length has been determined
	var arr1 = [3]int{1, 2, 3}
	arr2 := [2]string{"Rifaldi", "noviansyah"}

	// array with no length
	var arr3 = [...]int{1, 2, 4}
	arr4 := [...]string{"Rifaldi", "noviansyah"}

	// array with length but no data
	var arr5 = [5]int{}
	var arr6 = [5]string{}

	// array with length but no fuull
	var arr7 = [5]int{1, 2, 3}

	// array with intial index
	var arr8 = [5]int{0: 10, 2: 30}

	// change values array
	arr1[0] = 100

	// array slice
	// if we use slice, initial array  must be [] empty array
	arr9 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice1 := arr9[0:2]

	fmt.Println("array 1 = ", arr1)
	fmt.Println("array 2 = ", arr2)
	fmt.Println("array 3 = ", arr3)
	fmt.Println("array 4 = ", arr4)
	fmt.Println("array 5 = ", arr5)
	fmt.Println("array 6 = ", arr6)
	fmt.Println("array 7 = ", arr7)
	fmt.Println("array 8 = ", arr8)
	// akses array
	fmt.Println("array 1 index ke 0 = ", arr1[0])
	// array slice
	fmt.Printf("array slice = %v\n", myslice1)
	// knowing length array
	fmt.Printf("knowing length array = %d\n", len(myslice1))
	// knowing capacity array
	fmt.Printf("knowing capacity array = %d\n", cap(myslice1))
	// append for slice in the end of array
	arr9 = append(arr9, 20, 21)
	fmt.Printf("knowing append array = %v\n", arr9)

	// copy
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// original
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// create copy with only needed numbers
	// make function can also be used to create a slice
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
