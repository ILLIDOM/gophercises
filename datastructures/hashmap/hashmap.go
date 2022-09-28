package main

import "fmt"

/*
https://go.dev/blog/maps
*/

func main() {
	// initialization empty map (nil map)
	var map1 map[int]string // keys are of type int and values of type string
	// map1[1] = "value1"      // error because nil map
	fmt.Println(map1)

	// initialize map the correct way
	map2 := make(map[int]string)
	map2[99] = "value1"
	i := map2[99] // doesnt panic if key is not present; default value is returned
	fmt.Println(i)

	// add key
	map2[8] = "test"

	// delete key
	delete(map2, 99)

	// get key
	i, ok := map2[8] // ok is true if key is present else false
	if ok {
		// use value
		fmt.Println(i)
	}

	// loop key, values
	for key, value := range map2 {
		fmt.Println("Key:", key, "Value:", value)
	}

	// number of items in map
	n := len(map2)
	fmt.Printf("lenght of map2 is: %d\n", n)
}
