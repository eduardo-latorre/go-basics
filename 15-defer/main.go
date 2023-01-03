package main

import "fmt"

/*
Defer
- Are execute at the end of the program
- When multiple defer function get executed like LIFO (Last declared First executed)
- Ensure connections to DB and Brokers are closed
- Very useful to close http body request
*/

func main() {
	defer deferFunction() // This will be print last cause it's first defer
	defer fmt.Println("Third")
	fmt.Println("First")
	fmt.Println("Second")
	deferFunction2()
	fmt.Print("\n")
}

func deferFunction() {
	a := 0
	for a < 5 {
		fmt.Print(a)
		a++
	}
	fmt.Print("\n")
}

// This will revert printing instead
func deferFunction2() {
	a := 5
	for a < 10 {
		defer fmt.Print(a)
		a++
	}
}
