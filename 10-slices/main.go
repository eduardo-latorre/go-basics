package main

import (
	"fmt"
	"sort"
)

/*
Slices
- Slices doesn't have fixed sized and it's only one type
- Values can be repeated
- It's a slice of an array literally
- Does not have a fixed size
- It can be extended
*/
func main() {
	// Example - 01

	// Basic declaration, no length needed and needs to be instantiated using {}
	var basicSlice = []string{}
	fmt.Printf("Value type: %T\n", basicSlice)

	// Populated slice
	var jediMasters = []string{"Yoda", "Mace Window"}
	fmt.Println(jediMasters)

	// Appending the Slice
	jediMasters = append(jediMasters, "Obi Wan", "Anakin", "Yoda")
	fmt.Println(jediMasters)

	//Slicing a regular array
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}

	// Creating a slice from index 0 to 3, range are not inclusive so it won't count value from position 2
	// [1, 2]
	var slice []int = array1[0:2]
	fmt.Println(slice)

	//Diference between array and slices:
	//- slices and arrays have the same capacity but different length
	fmt.Printf("Array lenght: %v and capacity: %v\n", len(array1), cap(array1))
	fmt.Printf("Slice lenght: %v and capacity: %v\n", len(slice), cap(slice))

	// Example - 02 creating a slice from an implicit array
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

	// Example - 03 Using make to create a Slice, it's need the begging size of the Slice
	qualifications := make([]float64, 4)

	qualifications[0] = 4.4
	qualifications[1] = 4.1
	qualifications[2] = 6.9
	qualifications[3] = 5.6

	fmt.Println(qualifications)

	//As it can't be added manually a new value in the Slice it can be done by using Append
	qualifications = append(qualifications, 4.3, 5.3, 4.2, 4.8, 7.0)

	fmt.Println(qualifications)

	//EXERCISE - sort values from lower to higher
	lenght := len(qualifications)

	for i := 0; i < lenght; i++ {
		for j := i; j < lenght; j++ {
			if qualifications[i] > qualifications[j] {
				aux := qualifications[i]
				qualifications[i] = qualifications[j]
				qualifications[j] = aux
			}
		}
	}
	fmt.Println(qualifications)

	// Example - 04 Sorting Slices
	sizes := []float64{}
	sizes = append(sizes, 1.692, 1.691, 1.72, 1.54)
	fmt.Println("Sizes are sorted?", sort.Float64sAreSorted(sizes))
	sort.Float64s(sizes)
	fmt.Println(sizes)
}
