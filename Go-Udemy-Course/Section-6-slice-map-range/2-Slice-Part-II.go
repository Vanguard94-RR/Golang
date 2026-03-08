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
	//src[1] = "w"

	// extend a slice append new values to it
	src = append(src, "a", "b", "c")
	fmt.Println(src)

	// another way to create a slice with predefined length and capacity
	dest := make([]string, 6)
	fmt.Println(dest)

	// copy a slice to another slice
	copy(dest, src)
	fmt.Println(dest)

	//Since Slice ins mainly a pointer to underlying array, slicing a slice doesnot create a new array,
	// it just creates a new slice that points to the same underliying array.
	// So if we modify the original slice, the changes will be reflected in the new slice as well.

	// S1 := src[0:3] // this will create a new slice that points to the same underlying array as src,
	//  it will include the elements from index 0 to 2 (3 is exclusive)
	S1 := src[0:3]
	fmt.Println(S1)

	// S2 := src[:3] // this will create a new slice that points to the same underlying array as src,
	// it will include the elements from index 0 to 2 (3 is exclusive)
	S2 := src[:3]
	fmt.Println(S2)

	// S3 := src[3:] // this will create a new slice that points to the same underlying array as src,
	// it will include the elements from index 3 to the end of the slice
	S3 := src[3:]
	fmt.Println(S3)

	// S4 := src[:] // this will create a new slice that points to the same underlying array as src,
	// it will include all the elements of the slice
	S4 := src[:]
	fmt.Println(S4)

	// modify the original slice
	src = append(src, "o", "p", "q")
	fmt.Println(src)
	fmt.Println(S1)
	fmt.Println(S2)
	fmt.Println(S3)
	fmt.Println(S4)
}
