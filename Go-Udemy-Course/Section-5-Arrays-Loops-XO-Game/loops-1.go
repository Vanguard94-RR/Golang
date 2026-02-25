package main

import "fmt"

func main() {

	//for loops

	//
	for i := 0; i <= 5; i++ {
		fmt.Println("Num", i)
	}

	//for with condition only
	fmt.Println("for x <=3")

	x := 1

	for x <= 3 {
		fmt.Println("Num", x)
		x++
	}
	fmt.Print("x Value after loop", x)
}
