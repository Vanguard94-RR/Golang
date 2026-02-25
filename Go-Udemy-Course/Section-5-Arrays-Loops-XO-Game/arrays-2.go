package main

import "fmt"

func main() {

	var arr [6]string
	fmt.Println("", arr)
	// Arrays
	arr[0] = "Juan"
	arr[1] = "Maria"
	arr[2] = "Pedro"
	arr[3] = "Ana"
	arr[4] = "Luis"
	arr[5] = "Sofia"

	fmt.Println("String array:", arr)

	arr = [6]string{"Juan", "Maria", "Javier", "Ana", "Luis", "Sofia"}

	fmt.Println("Reasigned String array:", arr)

	// Array with inferred length
	var num = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Int array:", num)

	// Accessing array elements
	fmt.Println("4th number", num[3])

	// length of array
	fmt.Println("Length of num array:", len(num))
}
