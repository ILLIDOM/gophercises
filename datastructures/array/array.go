package main

import "fmt"

func main() {

	// Initialize Array
	var arr1 [5]int // declare array of size 5
	arr1[0] = 0

	arr2 := [5]int{1, 2, 3, 4, 5} // declare and init array with values
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3} // declare and init array (compiler determines size but is also fix)
	fmt.Println(arr3)
}
