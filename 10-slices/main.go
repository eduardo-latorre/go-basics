package main

import "fmt"

/*
Slices
- It's a slice of an array literally
- Does not have a fixed size
- It can be extended
*/
func main() {
	// Example - 01
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}

	// Creating a slice from index 0 to 3
	var slice []int = array1[0:2]
	fmt.Println(slice)

	//Diference between array and slices:
	//- slices and arrays have the same capacity but different length
	fmt.Printf("Array lenght: %v and capacity: %v\n", len(array1), cap(array1))
	fmt.Printf("Slice lenght: %v and capacity: %v\n", len(slice), cap(slice))

	// Example - 02, creating a slice from and implicit array
	var slice1 []int = []int{1, 2, 3, 4}

	// Then adding a new element to the slice and returning a new slice
	// Append add a new element and returns a new slice
	var newSlice1 []int = append(slice1, 5)
	fmt.Println(slice1)
	fmt.Println(newSlice1)

	// Creating a new slice and adding the returning value to the same var
	var slice2 []string
	slice2 = append(slice2, "Superman", "Wolverine")
	fmt.Println(slice2)
}
