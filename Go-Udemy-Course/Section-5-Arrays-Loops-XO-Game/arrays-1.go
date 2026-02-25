package main

import "fmt"

func main() {
	// define arrays

	var arr [6]string
	fmt.Println("", arr)

	var a [3]int
	fmt.Println("int array:", a)

	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println("String array:", arr)

	arr[2] = "Go"
	arr[3] = "Programming"
	arr[4] = "Language"
	arr[5] = "!"
	fmt.Println("String array:", arr)

	a[0] = 10
	a[1] = 20
	a[2] = 30
	fmt.Println("int array:", a)

}
