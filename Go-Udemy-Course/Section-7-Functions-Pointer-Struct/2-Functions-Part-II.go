package main

import (
	"errors"
	"fmt"
)

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

func sum(a int, b int) int {
	return a + b
}

// Function that returns multiple results.
func concat(a, b string) (string, int) {
	result := a + b
	return result, len(result)
}

// Funtion that returns nil and string or default string value
func returnError(flag bool) (string, error) {
	if flag {
		return "", errors.New("This is a go custom error")
	}
	return "This is a default string value", nil
}

func main() {
	Tittle("Functions Part II")

	Tittle("Multiple Values Returned in here")
	result, length := concat("Hello, ", "World!")
	fmt.Println("Concatenated string:", result)
	fmt.Println("Length of concatenated string:", length)

	Tittle("Error Handling in Go")
	value, err := returnError(false) // change to false to get default string value
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Value:", value)
	}

	// Funtions types in Go
	Tittle("Function types")
	// In Go, functions are first-class citizens, which means they can be treated like any other type.
	// You can assign a function to a variable, pass it as an argument to another function, or return it from a function.
	var x func(int, int) int
	x = sum
	fmt.Println("Sum of 10 and 20 using function variable:", x(10, 20))
}
