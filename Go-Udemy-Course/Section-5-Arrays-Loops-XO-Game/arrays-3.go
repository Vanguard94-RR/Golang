package main

import "fmt"

func main() {

	// 2D array
	XoBoard := [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	fmt.Println(XoBoard)
	XoBoard = [3][3]string{
		[3]string{"-", "-", "x"},
		[3]string{"-", "o", "-"},
		[3]string{"x", "-", "o"},
	}
	fmt.Println(XoBoard[0])
	fmt.Println(XoBoard[1])
	fmt.Println(XoBoard[2])

	XoBoard[0][1] = "x"
	fmt.Println(XoBoard[0])
	fmt.Println(XoBoard[1])
	fmt.Println(XoBoard[2])
}
