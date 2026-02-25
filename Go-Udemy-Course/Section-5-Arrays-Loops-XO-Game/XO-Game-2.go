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
		// Read row and column values from user input
		var row int
		fmt.Println("Please enter number 0,1,2")

		fmt.Scanln(&row)
		if row < 0 || row > 2 {
			fmt.Println("Invalid row number \n Please enter number 0,1,2")
			continue
		}
		var col int
		fmt.Println("Please enter number 0,1,2")

		fmt.Scanln(&col)
		if col < 0 || col > 2 {
			fmt.Println("Invalid column number \n Please enter number 0,1,2")
			continue
		}

		//Set the player's move on the board
		if board[row][col] == "" {
			board[row][col] = Player

		} else {
			// Cell is already occupied, prompt the player to choose another cell
			fmt.Println("Cell is already occupied", row, col, "value is", board[row][col], "\n Please choose another cell")
			continue
		}

		// Print the current state of the board
		fmt.Println("Current Board:")
		fmt.Println(board[0])
		fmt.Println(board[1])
		fmt.Println(board[2])

		// Check for a win condition
		for i := 0; i < 3; i++ {
			// Check rows
			if board[i][0] == Player && board[i][1] == Player && board[i][2] == Player ||
				board[0][i] == Player && board[1][i] == Player && board[2][i] == Player {
				fmt.Println("Player", Player, "wins!")
				return
			}

			// Check diagonals
			if board[0][0] == Player && board[1][1] == Player && board[2][2] == Player ||
				board[0][2] == Player && board[1][1] == Player && board[2][0] == Player {
				fmt.Println("Player", Player, "wins!")
				return
			}

			// Switch player
			if Player == "X" {
				Player = "O"
			} else {
				Player = "X"
			}
		}
	}
}
