package main

import "fmt"

func Tittle(text string) {
	dashtitle := ""
	for i := 0; i < (78-len(text))/2; i++ {
		dashtitle += "-"
	}
	fmt.Println(dashtitle + text + dashtitle)
}
