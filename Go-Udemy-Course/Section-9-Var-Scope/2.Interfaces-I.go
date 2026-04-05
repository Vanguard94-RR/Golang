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

const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Italic    = "\x1b[3m"
	Underline = "\033[4m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Blue      = "\033[34m"
	Yellow    = "\033[33m"
	Blink     = "\x1b[5m"
	Inverse   = "\x1b[7m"
	Gray      = "\x1b[37m"
	DarkGray  = "\x1b[30;1m"
)

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
	fmt.Println(Bold + left + Green + Italic + text + Reset + Bold + right + Reset)
}

func subTittle(text string) {
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
	fmt.Println(left + DarkGray + text + Reset + right)
}

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
	Juan := &human{}

	steps := []point{
		point{x: 1, y: 2},
		point{x: 3, y: 4},
		point{x: 4, y: 6},
	}
	err := move(Juan, steps)
	if err != nil {
		fmt.Println("Error moving:", err)
	} else {
		fmt.Println("Final position of Juan:", Juan.getPosition())
	}

	Juan.talk("Hello, I am Juan!")
}

func main() {
	Tittle("Interfaces in Go")
	subTittle("Defining and Implementing Interfaces")
	interfaces()
}
