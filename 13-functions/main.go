package main

import "fmt"

// Functions
// Cannot create emmbebed functions
// Returning values needs set the returining type, named as function signatures

func main() {

	// 01 - Calling a simple function
	greeter()

	// 02 - Using function passing values
	fmt.Println("Let's add two number!")

	var num1 int
	var num2 int

	fmt.Printf("Enter your first number: ")
	fmt.Scan(&num1)

	fmt.Printf("Enter your second number: ")
	fmt.Scan(&num2)

	fmt.Printf("The result of the add is: %v\n", adder(num1, num2))

	// 03 - Using a function with multiple values and multiple return
	var result, values int
	result, values = proAdder(2, 4, 5, 6, 3, 2, 5, 67, 3)
	fmt.Printf("Result of Pro Adder is %v of %v number added\n", result, values)
}

// Simple function
func greeter() {
	fmt.Println("Welcome to this functions app example in Golang!")
}

// Adding two numbers
func adder(num1 int, num2 int) int {
	return num1 + num2
}

// Adds multiple numbers, it expects a slice and returns result and amount of values added
func proAdder(values ...int) (int, int) {
	var result int
	for _, value := range values {
		result += value
	}
	return result, len(values)
}
