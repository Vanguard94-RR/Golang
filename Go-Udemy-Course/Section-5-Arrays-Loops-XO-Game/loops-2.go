package main

import "fmt"

func main() {

	separator := "------------------"

	//for loops

	//
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println("Num", i)
	// }

	// //for with condition only
	// fmt.Println("for x <=3")

	// x := 1

	// for x <= 3 {
	// 	fmt.Println("Num", x)
	// 	x++
	// }
	// fmt.Print("x Value after loop", x)

	//infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	//for with break and continue
	for a := 0; a <= 5; a++ {
		if a == 3 {
			continue
		}
		fmt.Println(a)

	}
	fmt.Println(separator)
	//break
	for {
		fmt.Println("This will run one time")
		fmt.Println(separator)
		break
	}
	fmt.Println(separator)

	//nested loops
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println("x,y", x, y)
		}
	}

}
