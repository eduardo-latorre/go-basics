package main

import "fmt"

func main() {

	// Simple if statement
	jedi := true

	if jedi {
		fmt.Println("The force is with you")
	} else {
		fmt.Println("You're surrounded by the dark side of the force")
	}

	// Another simple way is
	age := 30
	apprentice := age <= 20

	if apprentice {
		fmt.Println("You're just an apprentice")
	} else {
		fmt.Println("You're ready to be a Knight")
	}

	// Else if, or/and statement
	lightSaberColor := "Green"

	if lightSaberColor == "Red" {
		fmt.Println("You're a sith")
	} else if lightSaberColor == "Blue" || lightSaberColor == "Green" {
		fmt.Println("You're a jedi")
	} else {
		fmt.Println("You're a very special ")
	}

	// Unequal statement
	if lightSaberColor != "Blue" || lightSaberColor != "Green" && age > 20 {
		fmt.Println("You should be the lord sith we're looking for all these years")
	} else {
		fmt.Println("You should be a jedi master")
	}

}
