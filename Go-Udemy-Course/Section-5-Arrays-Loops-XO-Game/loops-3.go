package main

import "fmt"

func main() {

	separator := "------------------"

	//for with break
	fmt.Println(separator)

breakPoint:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				break breakPoint
			}
			fmt.Println("x,y", x, y)
		}
	}
	fmt.Println(separator)

	//for with
}
