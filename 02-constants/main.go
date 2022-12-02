package main

import "fmt"

// Basic declaration using standard for constants, using capital letters
const SUPERHERO = "Superman"

func main() {
	// Mutiple declaration
	const SUPER_POWER1, SUPER_POWER2, SUPER_POWER3, SUPER_POWER4 = "Fly", "Steel Sking", "Super Speed", "Frezzing Breath"

	//Mutiple declaration using brakets
	const (
		VILLAIN1 = "Lex Luthor"
		VILLAIN2 = "Dooms Day"
		VILLAIN3 = "Black Adams"
	)

	fmt.Println(SUPERHERO)
	fmt.Println(SUPER_POWER1, SUPER_POWER2, SUPER_POWER3, SUPER_POWER4)
	fmt.Println(VILLAIN1, VILLAIN2, VILLAIN3)
}
