package main

import "fmt"

/*
https://go.dev/blog/slices-intro

Slice contains
	- pointer to underlying array
	- capacity (max elements possible)
	- length (currently inserted elements)

if capacity is full, a new underlying array is created and all eements are copied

passing a slice as an funciton argument only the pointer is copied (shallow copy). the undelying array is the same

slicing creates a new slice (pointer). the underlying array stays the same - what changes is the first elements the pointer points to and the lenght
*/

func main() {
	// initialize slice
	var slice1 []string // declare slice

	slice1 = append(slice1, "abc")

	slice2 := []string{"abc", "abd"} // declare and init slice with values
	slice2[1] = "bbb"

	slice3 := make([]byte, 5, 5)
	fmt.Println(slice3)

	// Append to array using copy
	slice1 = append(slice1, "new element")

	// Deep Copy Array using append
	src := []string{"e1", "e2", "e3"}
	dst := make([]string, len(src)) // create a new slice with empty array of lenght src
	copy(dst, src)                  // copy min(len(dst), len(src)) elements from src to dst

	// Deep copy 2 using append
	src1 := []int{1, 2, 3, 4}
	var dst1 []int
	dst1 = append(dst1, src1...) // append all elements inside src1 to dst1

	// slicing operator
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:4]

	fmt.Println("Original Array: 1,2,3,4,5")
	// address of array is the first element of the slice
	fmt.Printf("Address of array is: %p, length is %d, capacity is %d\n", &s1, len(s1), cap(s1))
	fmt.Println("After slicing operator [1:4]")
	// address of array is the first element of the slice
	fmt.Printf("Address of array is: %p, length is %d, capacity is %d\n", &s2, len(s2), cap(s2))
}
