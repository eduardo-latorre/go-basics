package main

import "fmt"

// Package variables - Can't shorthand declaration
var superHero string = "Superman"

func main() {

	// Variables
	// int, string, float32, float64, bool

	// 01- Standard declaration
	var name string = "Clark"

	fmt.Println(name)

	// 02- Shorthand variable - Only used inside methods
	lastName := "Kent"
	age := 33

	fmt.Println(lastName, age)
	fmt.Printf("Lastname var type: %T, and Age var type: %T\n", lastName, age)

	// 3- Declared but not used
	var nationality string

	nationality = "american"
	fmt.Println(nationality)

	// 4- Multiple declaration
	var height, weight = 1.85, 95.3
	fmt.Printf("heigth: %v and weight: %v\n", height, weight)
}
