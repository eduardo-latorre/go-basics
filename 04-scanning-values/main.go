package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	var lastName string
	var birthYear int

	fmt.Println("Welcome to this small app!")

	time.Sleep(2 * time.Second)

	fmt.Println("What's your name?")
	fmt.Scan(&name)

	time.Sleep(1 * time.Second)

	fmt.Println("What's your Lastname?")
	fmt.Scan(&lastName)

	time.Sleep(1 * time.Second)

	fmt.Println("What's your birth year?")
	fmt.Scan(&birthYear)

	time.Sleep(2 * time.Second)

	fmt.Printf("Welcome %v %v, your age this year is or will be %v\n", name, lastName, 2022-birthYear)

}
