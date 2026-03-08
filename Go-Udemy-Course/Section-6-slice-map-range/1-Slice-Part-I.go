package main

import "fmt"

func main() {
	// Slice is a reference type, it does not store any data, it just points to an underlying array.
	// The length of a slice is the number of elements in the slice, while the capacity is the number of elements in the underlying array.
	// When you append to a slice and it exceeds its capacity, a new underlying array is created and the existing elements are copied to it.

	// Empty slice
	keywords := []string{}
	fmt.Println(keywords)

	// Slice with values
	src := []string{"x", "y", "z"}
	fmt.Println(src)

	// get/set values by index
	fmt.Println(src[0]) // x
	src[1] = "y"

	// extend a slice append new values to it
	src = append(src, "a", "b", "c")
	fmt.Println(src)
}
