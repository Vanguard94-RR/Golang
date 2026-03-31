package main

import "fmt"

func Tittle(text string) {
	leftDashes := (78 - len(text)) / 2
	rightDashes := 78 - len(text) - leftDashes

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
