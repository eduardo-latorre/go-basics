package main

import "fmt"

/*
Arrays
- Always has a fixed size
- Always have the same data type
*/
func main() {
	// 01 - Example
	// Simple way to create an array
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("Array1 values: %v\n", array1)

	// Length and capacity are always going to be pretty much the same
	fmt.Printf("Array length: %v\n", len(array1))
	fmt.Printf("Array capacity: %v\n", cap(array1))

	// 02 - Example
	// Declaring it, but not using it
	var array2 [5]int

	array2[0] = 1
	array2[1] = 10
	array2[2] = 100
	array2[3] = 1000
	array2[4] = 10000

	fmt.Printf("Array2 values: %v\n", array2)

	// 03 - Example
	// Using other data types
	var array3 [3]string = [3]string{"Superman", "The Batman", "Flash"}

	fmt.Printf("Array3 values: %v\n", array3)
}
