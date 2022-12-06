package main

/*
	This is used when multiple if's
	keywords:
	- break: it's assumed by Go, it's not mandatory
	- fallthrough: allow to pass to next case
*/
import (
	"fmt"
	"runtime"
)

func main() {

	// 01 - Example
	operativeSystem := runtime.GOOS
	fmt.Println(operativeSystem)

	switch operativeSystem {
	case "darwin":
		fmt.Println("This program is running on Darwin")
		fallthrough //Allows to pass to next case "linux"
	case "linux":
		fmt.Println("This program is running on Linux")
	case "windows":
		fmt.Println("This program is running on Windows")
	default:
		fmt.Println("Not know OP")
	}

	// 02 - Example
	age := 2

	//When value returns a bool, it's not necesary to use a var after switch statement.
	//But in this case you can use "true" to be more explicit
	switch {
	case age >= 18:
		fmt.Println("You're an adult")
	case age > 12 && age < 18:
		fmt.Println("You're an adult")
	default:
		fmt.Println("You're a kid")
	}

}
