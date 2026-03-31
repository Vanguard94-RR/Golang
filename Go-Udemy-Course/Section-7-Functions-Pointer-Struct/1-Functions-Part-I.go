package main

import "fmt"

func Tittle(text string) {
	dashtitle := ""
	for i := 0; i < (78-len(text))/2; i++ {
		dashtitle += "-"
	}
	fmt.Println(dashtitle + text + dashtitle)
}

// A function is a reusable block of code that performs a specific task.
// Functions allow you to break down your code into smaller, more manageable pieces,
// and they can be called multiple times throughout your program.
// In Go, you can define a function using the func keyword, followed by the function name,
// a list of parameters (if any), and the function body enclosed in curly braces.
// Functions can also return values.

// This function takes a string parameter and prints a message, but it does not return any value.
func empty(s string) {
	Tittle("Return Nothing")
	fmt.Println("Hello,", s, "this function doesn't return a value")

}

// This funtion take as many parameters as you need, but does not return any value.
// It ptints a message with the parameters length and the parameters itself.

func manyParams(params ...interface{}) {
	Tittle("Many Parameters")
	fmt.Println("Number of parameters:", len(params))
	fmt.Println("Parameters:", params)
}

// In Go, you can also define functions that return values. To specify the return type of a function,
// you include it after the parameter list. For example, a function that takes two integers and returns
// their sum would be defined like this:

func sum(a int, b int) int {
	return a + b
}

func main() {
	Tittle("Functions Part I")

	empty("User")
	manyParams("Hello", 42, true, 3.14, []string{"Go", "Python", "Java"})

	fmt.Println("Sum of 3 and 5 is:", sum(3, 5))
}
