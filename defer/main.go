package main

import "fmt"

/*
Defer
- Ensure connections to DB and Brokers are closed
- Very useful to close http body request
- Are execute at the end of the program
- When multiple defer function get executed like LIFO (Last declared First executed)
*/

type engineer struct {
	name string
}

func (e *engineer) updateName(name string) {
	e.name = name
}

func deferFunction(e *engineer) {
	name := "Eduardo Latorre"
	defer e.updateName(name)
	fmt.Println("I think this is going to be the last print")

	// In this case Go won't change the name value even if it's executed last
	name = "Javier Latorre"
}

func main() {
	fmt.Println("**** Defer tutorial ****")

	eduardo := &engineer{
		name: "Eduardo",
	}

	fmt.Printf("%+v\n", eduardo)
	deferFunction(eduardo)
	fmt.Printf("%+v\n", eduardo)
}

// Second example **********

// func deferFunction() (x int) {
// 	defer func() {
// 		x = 16 // Print this value since is passed last
// 	}()
// 	x = 5
// 	return
// }

// func main() {
// 	fmt.Println("**** Defer tutorial ****")
// 	fmt.Println(deferFunction())
// }

// First Example **********

// func main() {
// 	fmt.Println("**** Defer tutorial ****")

// 	defer func() {
// 		fmt.Println("First defer function - Executed Last")
// 	}()

// 	defer func() {
// 		fmt.Println("Last defer function - Executed First")
// 	}()

// 	fmt.Println("**** End of the program ****")
// }
