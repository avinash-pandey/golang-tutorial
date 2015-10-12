package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	//there are five basic types in golang
	//boolean
	var boolean bool = true

	//string
	var stringVariable string = "SomeString"

	//integer with sighned and unsighned types
	//int defaults to the word size of your operating system
	var integerVariable int32 = 444

	//there are two types of float values

	var floatVariableOne float32 = 44.0

	var floatVariableTwo float64 = 55.5555

	//the last type is the complex type defined in the math package

	var complexVariable complex128 = cmplx.Sqrt(-5)

	fmt.Println(boolean, stringVariable, integerVariable, floatVariableOne, floatVariableTwo, complexVariable)

}
