package main

import (
	"errors"
	"fmt"
	"os"
)

// Errors
// - They don't break the app as panic
// - Custom error messages shouldn't capitalized

// Creating a custom error
func myCustomError() error {
	return errors.New("this is my custom error and shouldn't be capitalized")
}

// Diving two numbers, first example
func divideByZero(a, x float64) (float64, error) {
	if x == 0 {
		return float64(-1), errors.New("can't divide by 0")
	} else {
		return float64(a) / float64(x), nil
	}
}

// Creating my own custom error
type MyCustomError struct {
	msg string
	err error
}

func (myCustomError *MyCustomError) Error() string {
	return myCustomError.msg
}

func divideByZero2(a, x float64) (float64, error) {
	if x == 0 {
		return float64(-1), &MyCustomError{
			msg: "this is just a custom message",
			err: errors.New("this is the real error"),
		}
	}
	return float64(a) / float64(x), nil
}

func main() {

	// Example 01 - Creating a simple error
	fmt.Println(myCustomError())

	// Example 02 - Deviding by 0
	result, resultErr := divideByZero(10, 0)

	if resultErr == nil {
		fmt.Println("The result is: ", result)
	} else {
		fmt.Println("error is: ", resultErr)
	}

	// Example 03 - Reading file

	// INTENDED ERROR
	file, fileErr := os.ReadFile("file-whatever")

	// This is how errors are caught
	if fileErr == nil {
		fmt.Println(string(file))
	} else {
		fmt.Println(fileErr)
	}

	fmt.Println("The program continues...")

	// Example 04 - Creating my custom error
	result2, resultErr2 := divideByZero2(144, 0)

	if resultErr2 != nil {
		var divErr *MyCustomError
		errors.As(resultErr2, &divErr)
		fmt.Println("This is not matematically valid", resultErr2.Error())
	}
	fmt.Println("Result 2 from division: ", result2)
}
