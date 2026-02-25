package main

import "fmt"

func main() {
	//XO Game
	// initialize the board
	board := [3][3]string{}
	Player := "X"

	for {

		fmt.Println("Player:", Player)

		//read board values
		var row int
		var col int

		fmt.Println("Please enter number 0,1,2")
		fmt.Scanln(&row)

		// Read row and column values from user input
		if row < 0 || row > 2 {
			fmt.Println("Invalid row number \n Please enter number 0,1,2")
			continue
		}
		if col < 0 || col > 2 {
			fmt.Println("Invalid column number \n Please enter number 0,1,2")
			continue
		}

	}
}
