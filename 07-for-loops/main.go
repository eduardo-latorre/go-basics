package main

import (
	"fmt"
)

func main() {

	// 01 - Basic example
	loops := 5

	fmt.Printf("This will iterate %v times\n", loops)
	for i := 0; i < loops; i++ {
		fmt.Println(i)
	}

	// 02 - Iterating an array
	climbingGear := []string{"Rope", "Harness", "Quickdraws", "Grigri", "Climbing shoes"}

	for i := 0; i < len(climbingGear); i++ {
		fmt.Printf("Position %v gear %v\n", i, climbingGear[i])
	}

	// 03 - Shorter sintax of using loops
	for i := range climbingGear {
		fmt.Println(climbingGear[i])
	}

	// 04 - Using for loop as while
	loopWhile := 1

	for loopWhile < 5 {
		loopWhile += loopWhile
		fmt.Println(loopWhile)
	}

	// 05 - Using break in the loop
	var numbers [5]int
	numbers = [5]int{23, 24, 25, 26, 27}

	// Shoud stop until it's equal 25
	for i := range numbers {
		if numbers[i] == 25 {
			break
		}
		fmt.Println("Number hit:", numbers[i])
	}

	// 06 - Using continue in the loop
	var vowels [5]string
	vowels = [5]string{"a", "e", "i", "o", "u"}

	// Shoud not print vowel e but it still continues
	for i := range vowels {
		if vowels[i] == "e" {
			continue
		}
		fmt.Println("Number hit:", vowels[i])
	}

	// 07 - Using continue goto (labels)
	topWorldClimber := [5]string{"Adam", "Chris", "Alex", "Daniel", "Emil"}

	for i := range topWorldClimber {
		if topWorldClimber[i] == "Alex" {
			goto findClimber
		}
		fmt.Println(topWorldClimber[i])
	}

findClimber:
	fmt.Println("We've found him")

	// 08 - For each
	fmt.Println("Climbing spots in Chile")
	climbingSpots := [5]string{"Palestras", "La Mina", "Los Manyos", "Piedra Romel", "Las Chilcas"}

	// for requires the index and object, in this case it's not using index
	for _, cliclimbingSpotObject := range climbingSpots {
		fmt.Println(cliclimbingSpotObject)
	}
}
