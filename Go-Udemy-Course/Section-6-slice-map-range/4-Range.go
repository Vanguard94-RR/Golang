package main

import "fmt"

func separador() {
	fmt.Println("--------------------------------------------------")
}

func main() {
	// Range is a funcion that allows us to iterate over a slice, array, map, or string.
	// It returns two values: the index and the value of the element at that index.

	// Iterate over an array
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	separador()
	// Iterate over a slice
	slice := []string{"a", "b", "c", "d", "e"}
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	separador()
	// Iterate over a map
	m := map[string]int{
		"Alice":   10,
		"Bob":     20,
		"Charlie": 30,
	}
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	separador()
	// Iterate over a string
	str := "Hello, World!"
	for i := range str {
		fmt.Printf("Index: %d, Value: %c\n", i, str[i])
	}

}
