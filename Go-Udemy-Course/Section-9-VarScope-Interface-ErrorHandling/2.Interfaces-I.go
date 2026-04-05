// Interfaces are types that define a set of method signatures. They are used to specify behavior that a struct must implement
// A type implements an interface by implementing its methods. There is no explicit declaration of intent,
// no "implements" keyword. Implicit interfaces are satisfied by any type that has the required methods.
// Interfaces are a powerful way to achieve polymorphism in Go, allowing you to write flexible and reusable code.
// They enable you to define behavior that can be shared across different types,
// without requiring those types to be related through inheritance.

package main

import (
	"errors"
	"fmt"
)

type walker interface {
	walk(p point) error
	getPosition() point
}

type talker interface {
	talk(s string)
}

type point struct {
	x int
	y int
}

// human struct that implements both walker and talker interfaces
type human struct {
	name     string
	position point
	speed    int
}

// Implementing the walker interface for human struct
func (h *human) walk(p point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("Invalid point: coordinates cannot be negative")
	}
	h.position = p
	fmt.Println("human walked to", h.position)
	return nil
}

// Implementing the getPosition method for human struct
func (h *human) getPosition() point {
	return h.position
}

// Implementing the talker interface for human struct
func (h *human) talk(s string) {
	fmt.Println("human says:", s)
}

func move(w walker, points []point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func interfaces() {
	PrintTitle("Interfaces Example")

	Juan := &human{}

	steps := []point{
		point{x: 1, y: 2},
		point{x: 3, y: 4},
		point{x: 4, y: 6},
	}

	PrintSubTitle("Moving Juan through points")
	err := move(Juan, steps)
	if err != nil {
		PrintError("Error moving: " + err.Error())
	} else {
		PrintSuccess("Final position of Juan: " + fmt.Sprintf("%v", Juan.getPosition()))
	}

	Juan.talk("Hello, I am Juan!")
}

func main() {
	PrintTitle("Interfaces in Go")
	PrintSubTitle("Defining and Implementing Interfaces")

	interfaces()

	PrintRoundBox("Demonstration Complete!")
}
