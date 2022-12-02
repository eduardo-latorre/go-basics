package main

import "fmt"

// Pointers

// Usage of & and *
// & = Pointer - Gets de address location in RAM
// * = Reference and Dereference operator - Gets the value from RAM location

func main() {
	// Print pointer address
	x := 7
	fmt.Printf("Var x value: %v and memory address: %v\n", x, &x)

	// Storing the pointer to another var
	y := &x
	fmt.Printf("Var y value from x address %v\n", y)

	// Changing the value from reference
	*y = 10 // Dereferenced
	fmt.Printf("Var x value changed by reference %v\n", x)

	// Example 02 - Using a pointer var
	actor := "Luke Skywalker"
	var pointer *string = &actor // Var expecting pointer address and passing pointer form actor var
	fmt.Printf("Pointer address: %v, reference value: %v", pointer, *pointer)

	// Example 03 - Passing the pointer to a function and changing the value from memory address
	movie := "Star Wars"
	fmt.Printf("Initial movie: %v\n", movie)
	changePointerValue(&movie)
	fmt.Printf("Changed movie value from pointe: %v\n", movie)
}

// This is expecting the pointer from the argument
func changePointerValue(movie *string) { // Expecting a pointer from string type
	*movie = "Harry Potter" // * Gets the value from the given address(pointer)
}
