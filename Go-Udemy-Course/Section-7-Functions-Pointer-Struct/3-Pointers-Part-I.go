// In Golang Pointers help us to work with memory addresses directly.
// A pointer is a variable that holds the memory address of another variable.
// Pointers allow us to manipulate data in memory, which can be useful for optimizing
// performance and managing resources efficiently.
// In Go, you can declare a pointer using the * operator.
// For example, to declare a pointer to an integer, you would write:
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

// Function that takes a pointer as an argument and modifies the value at that address
func modifyValue(p *int) {
	*p = 200 // Modifying the value at the address pointed to by p
}

func main() {

	Tittle("Pointers Part I")
	Tittle("Pointers Example I")
	var x int = 10
	var p *int = &x // p is a pointer to an integer, and it holds the address of x

	fmt.Println("Value of x:", x)          // Output: Value of x: 10
	fmt.Println("Address of x:", &x)       // Output: Address of x: 0xc0000140a8 (example)
	fmt.Println("Value of p:", p)          // Output: Value of p: 0xc0000140a8 (same as address of x)
	fmt.Println("Value at address p:", *p) // Output: Value at address p: 10 (dereferencing p)

	Tittle("Pointers Example II")

	name := "Alice"
	fmt.Println("Name:", name)             // Output: Name: Alice
	fmt.Println("Address of name:", &name) // Output: Address of name: 0xc000010220 (example)

	var namePtr *string = &name                        // namePtr is a pointer to a string, and it holds the address of name
	fmt.Println("Value of namePtr:", namePtr)          // Output: Value of namePtr: 0xc000010220 (same as address of name)
	fmt.Println("Value at address namePtr:", *namePtr) // Output: Value at address namePtr: Alice (dereferencing namePtr)

	// Modifying the value of name through the pointer
	*namePtr = "Bob"
	fmt.Println("Modified Name:", name) // Output: Modified Name: Bob (name is modified through the pointer)

	Tittle("Pointers Example III")

	// Creating a pointer to a new variable
	var num int = 42
	var numPtr *int = &num

	fmt.Println("Original value of num:", num)
	fmt.Println("Value through pointer:", *numPtr)

	// Modifying value through pointer
	*numPtr = 100
	fmt.Println("Modified value of num:", num)
	fmt.Println("Value through pointer after modification:", *numPtr)

	Tittle("Zero Value of Pointers")

	// Zero value of a pointer is nil
	var nilPtr *int
	fmt.Println("Zero value of pointer:", nilPtr)

	if nilPtr == nil {
		fmt.Println("Pointer is nil")
	}

	// Pass pointer to a function
	Tittle("Passing Pointers to Functions")

	var value int = 50
	fmt.Println("Original value:", value) // Output: Original value: 50

	modifyValue(&value) // Passing the address of value to the function

	fmt.Println("Modified value:", value) // Output: Modified value: 200 (value is modified through the pointer)
}
