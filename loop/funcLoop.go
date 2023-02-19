package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	arr := [2]string{"Rifaldi", "noviansyah"}

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// for idx, val := range arr {
	// 	fmt.Println(idx, val)
	// }

	for _, val := range arr {
		fmt.Println(val)
	}

}
