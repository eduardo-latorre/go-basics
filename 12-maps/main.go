package main

import "fmt"

// Maps
// - Key-value-pair, uses keys to store values
// - Don't store values in order
// Maps vs Arrays
// - To look for an element into an array you need to pass by all elements and it's slower
// - Maps just find the element by the index
// - Maps keeps time-complexity simpler it's o-1

func main() {

	// 01 - Basic declaration
	var lightSabers map[string]int = map[string]int{
		"blue":   4,
		"green":  3,
		"purple": 1,
	}

	fmt.Println(lightSabers)
	fmt.Println(lightSabers["green"])

	//Shorter declaration of maps and most used
	spaceships := make(map[string]string)

	spaceships["empire"] = "Deathstart"
	spaceships["rebels"] = "xwings"

	fmt.Println(spaceships)

	//Changing the value
	lightSabers["blue"] = 15

	//Adding new value
	lightSabers["yellow"] = 2

	fmt.Println(lightSabers)

	//Deleting values
	delete(lightSabers, "purple")
	fmt.Println(lightSabers)

	// 02 - Find an element inside the Map

	//val is the value of the map and ok is true or false whether exists or not
	val, ok := lightSabers["red"]
	fmt.Printf("%v %v\n", val, ok)

	// 03 - Map's lenght
	fmt.Println(len(lightSabers))

	//04 - Looping a map
	for _, v := range spaceships {
		fmt.Println(v)
	}
}
