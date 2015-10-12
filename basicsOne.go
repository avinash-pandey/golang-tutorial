package main

import "fmt"

func main() {
	//This is the most basic way to print something in go lang
	fmt.Println("Hello, avinash")

	//This is one way to declare variables in golang
	var someVariable string = "My name is avinash pandey"

	//In this way go lang can infer the type of the variable
	secondVariable := 5

	//All uninitialised variables in go have zero values
	var intVariable int
	var stringVariable string
	var floatVariable float32

	//Unused variables won't allow to get the code compile

	fmt.Println(someVariable, secondVariable, intVariable, stringVariable, floatVariable)

	//In order to use constants in go you need to use const keyword

	const n = 55

	fmt.Println(n)

	//this line would not compile because you cannot change a constant
	// n = 88
}
