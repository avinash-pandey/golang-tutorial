package main

//Go just has one loop which is the for loop
import (
	"fmt"
)

func main() {
	//The most familiar for loop .This won't work without the curly braces around the for loop body
	for i := 1; i < 10; i++ {
		fmt.Println("Hey just a try to check whether it works or not")
	}

	//just like the for loop in java or c you can leave any of the three parts of the loop
	for j := 1; ; j++ {
		fmt.Println("Second loop")
		break
	}

	//you can also use the for loop as the while loop
	iterator := 1
	for iterator < 5 {
		fmt.Println(iterator)
		iterator++
	}
}
