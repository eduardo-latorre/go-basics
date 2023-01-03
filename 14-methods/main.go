package main

import "fmt"

// Methods
// Belongs to a specific struct

type Jedi struct {
	name       string
	age        uint
	lightSaber string
	padawan    bool
}

// It's a best practice to use pointers because it uses the object intanciated directly from main function
// If I did otherwise, I would be passing a copy instead and print it
func (j *Jedi) getJedi() {
	fmt.Println(j)
}

func main() {
	newJedi := Jedi{
		name:       "Anakin",
		age:        13,
		lightSaber: "blue",
		padawan:    true,
	}

	newJedi2 := Jedi{
		name:       "Luke",
		age:        20,
		lightSaber: "blue",
		padawan:    true,
	}

	// As it using pointers, both should print the same value
	newJedi.getJedi()
	fmt.Println(newJedi)

	newJedi2.getJedi()
	fmt.Println(newJedi2)
}
