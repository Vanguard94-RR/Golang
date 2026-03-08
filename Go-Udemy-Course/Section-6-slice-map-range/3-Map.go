package main

import "fmt"

func main() {
	// A map is a collection of key-value pairs, where each key is unique and maps to a value.
	// Maps are implemented as hash tables, which provide fast access to values based on their keys.

	// init map with initial values
	m := map[string]bool{
		"a": true,
		"b": false,
		"c": true,
	}
	fmt.Println(m)

	// init map, with make function
	players := make(map[string]int)

	fmt.Println(players)

	//add values to map
	players["Alice"] = 10
	players["Bob"] = 20
	players["Charlie"] = 30

	fmt.Println(players)

	// list all keys in the map
	for name := range players {
		fmt.Println("Player:", name)
	}

	// Print  value from map by key
	fmt.Println("Alice's score:", players["Alice"])     // 10
	fmt.Println("Bob's score:", players["Bob"])         // 20
	fmt.Println("Charlie's score:", players["Charlie"]) // 30

	// delete a key-value pair from the map
	delete(players, "Bob")
	fmt.Println(players)

	// check if a key exists in the map
	score, ok := players["Bob"]
	if ok {
		fmt.Println("Bob's score:", score)
	} else {
		fmt.Println("Bob not found in the map")
	}

	// check if a key exists in the map
	score, ok = players["Alice"]
	if ok {
		fmt.Println("Alice's score:", score)
	} else {
		fmt.Println("Alice not found in the map")
	}

	// check if a key exists in the map
	score, ok = players["Charlie"]
	if ok {
		fmt.Println("Charlie's score:", score)
	} else {
		fmt.Println("Charlie not found in the map")
	}

	//append new key-value pair to the map
	players["Dave"] = 40
	players["Eve"] = 50
	players["Frank"] = 60
	fmt.Println(players)

	// create a map of slices

	grades := make(map[string][]int)
	grades["Alice"] = []int{90, 85, 92}
	grades["Bob"] = []int{80, 75, 88}
	grades["Charlie"] = []int{95, 92, 98}

	fmt.Println(grades)

}
