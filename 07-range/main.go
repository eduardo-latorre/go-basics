package main

import "fmt"

/*
Range
- is a better way to loop slices
*/
func main() {
	var slice []int = []int{45, 122, 23, 64, 55, 61, 70}

	//For loop - tradicional way
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Using range
	for index, element := range slice {
		fmt.Printf("Index: %v, Element: %v\n", index, element)
	}

	// Only using element without index
	for _, element := range slice {
		fmt.Printf("Element: %v\n", element)
	}

	// Problem, print duplicate elements
	fmt.Println("\n----- Exercise -----")
	slice2 := []int{1, 3, 4, 56, 7, 12, 4, 6, 1, 3}

	for i, elements1 := range slice2 {
		for j, elements2 := range slice2 {
			if elements1 == elements2 && i < j {
				fmt.Printf("Repeated value: %v\n", elements1)
			}
		}
	}
}
