package main

import "fmt"

func Tittle(text string) {
	leftDashes := (80 - len(text)) / 2
	rightDashes := 80 - len(text) - leftDashes

	left := ""
	for i := 0; i < leftDashes; i++ {
		left += "-"
	}
	right := ""
	for i := 0; i < rightDashes; i++ {
		right += "-"
	}
	fmt.Println(left + text + right)
}
