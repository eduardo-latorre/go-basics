package main

import "fmt"

// Structs are multi-type value format
// Are veru similar to Java classes
type engineer struct {
	name    string
	age     uint
	project project
}

type project struct {
	projectName  string
	priority     string
	technologies []string // slice
}

func main() {
	fmt.Println("**** Struct tutorial ****")

	newEngineer := engineer{
		name: "Eduardo",
		age:  26,
		project: project{
			projectName:  "Learning Golang",
			priority:     "High",
			technologies: []string{"Golang", "Ansible"},
		},
	}

	fmt.Printf("%+v\n", newEngineer) // %+v verbose description
}
