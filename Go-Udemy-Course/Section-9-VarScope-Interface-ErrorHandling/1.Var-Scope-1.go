// Variable scope in Go is the definition of where a variable can be accessed within the code. In Go,
// there are two main types of variable scope: package-level scope and function-level scope.

// Package-level scope means that a variable is accessible throughout the entire package, including all functions
// and files within that package. A variable declared at the package level can be accessed and modified by any
// function whithin the same package. Example:
/* var packageVar = "This is a package-level variable"

func anotherFunction() {
	fmt.Println("Accessing package variable from another function:", packageVar)
	// This demonstrates that packageVar can be accessed from any function in the same package
} */

// Function-level scope means that a variable is only accessible within the function where it is declared.
// Variables declared inside a function cannot be accessed from outside that function. Example:
/* func functionScopeExample() {
	var functionVar = "This is a function-level variable"
	fmt.Println(functionVar)
	// This will print "This is a function-level variable"
}	 */

// Block-level scope is a type of variable scope that is specific to a block of code, such as a loop or a conditional statement.
// Variables declared within a block are only accessible within that block. Example:
/* func blockScopeExample() {
	if true {
		var blockVar = "This is a block-level variable"
		fmt.Println(blockVar)
		// This will print "This is a block-level variable"
	}
	// fmt.Println(blockVar) // This will cause an error because blockVar is not accessible outside the if block
}
*/
// Global scope is a type of variable scope that is accessible throughout the entire program.
// Variables declared at the global scope can be accessed and modified by any function in the program. Example:
/* var globalVar = "This is a global variable"

func globalScopeExample() {
	fmt.Println(globalVar)
	// This will print "This is a global variable"
} */

// Local scope is a type of variable scope that is specific to a function or a block of code.
// Variables declared within a function or a block are only accessible within that function or block. Example:
/* func localScopeExample() {
	var localVar = "This is a local variable"
	fmt.Println(localVar)
	// This will print "This is a local variable"
	// fmt.Println(globalVar) // This will cause an error because globalVar is not accessible within this function
} */

// In Go, variables declared at the package level are accessible throughout the entire package, while variables
// declared within a function or a block are only accessible within that function or block.
// Understanding variable scope is important for writing clean and maintainable code in Go.

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

var name = "Juan" // This variable has package-level scope, it can be accessed from any function within this package

func main() {
	Tittle("Variable Scope in Go")
	Tittle("Package-Level")
	fmt.Println("Hello, my name is", name, "this variable has package-level scope, it can be accessed from any function within this package")

	Tittle("Function-Level")
	a := 10 // This variable has function-level scope, it can only be accessed within the main function
	fmt.Println("The value of a is:", a, "function-level scope means that a variable is only accessible within the function where it is declared.")

	Tittle("Block-Level")
	if a > 5 {
		b := 20 // This variable has block-level scope, it can only be accessed within this if block
		fmt.Println("The value of b is:", b, "block-level scope means that a variable is only accessible within the block where it is declared.")
	}
	// fmt.Println(b) // This will cause an error because b is not accessible outside the if block

	for i := 0; i < 3; i++ {
		c := i * 2 // This variable has block-level scope, it can only be accessed within this for loop
		fmt.Println("The value of c is:", c, "block-level scope means that a variable is only accessible within the block where it is declared.")
	}
	// fmt.Println(c) // This will cause an error because c is not accessible outside the for loop

	// Local scope example
	func() {
		Tittle("Local Scope")
		d := 30 // This variable has local scope, it can only be accessed within this anonymous function
		fmt.Println("The value of d is:", d, "Local scope means that a variable is only accessible within the function or block where it is declared.")
	}()
	// fmt.Println(d) // This will cause an error because d is not accessible outside the anonymous function

	// Demonstrating variable shadowing - when a variable in an inner scope has the same name as a variable in an outer scope
	shadowExample()
}

func shadowExample() {
	Tittle("Variable Shadowing")
	Tittle("Function Scope")
	name := "Outer scope name" // This variable is in the function scope
	fmt.Println("Before if block:", name, "this variable is in the function scope")

	if true {
		name := "Inner scope name" // This shadows the outer 'name' variable
		fmt.Println("Inside if block:", name, "this variable is in the inner scope and shadows the outer 'name' variable")
	}

	fmt.Println("After if block:", name, "back to the outer scope variable")

	// Demonstrating that the package-level 'name' variable is still accessible
	fmt.Println("Package-level name variable:", name, "this is the package-level variable that can be accessed from any function within this package")
}

// Function to demonstrate accessing package-level variables from different functions
func anotherFunction() {
	Tittle("Accessing Package-Level Variable from Another Function")
	fmt.Println("Accessing package-level variable 'name' from anotherFunction:", name)

	// We can also modify the package-level variable
	originalName := name
	name = "Modified from anotherFunction"
	fmt.Println("Modified package-level name:", name, "this shows that we can modify the package-level variable from any function within the same package")

	// Restore original value
	name = originalName
	fmt.Println("Restored package-level name:", name, "restoring the original value of the package-level variable")

}
