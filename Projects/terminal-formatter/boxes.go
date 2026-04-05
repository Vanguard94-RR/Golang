package main

import (
	"fmt"
	"strings"
)

// BoxStyle defines how a box is drawn
type BoxStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
	Color       string
}

// Box styles
var (
	SimpleBoxStyle = BoxStyle{
		TopLeft:     "+",
		TopRight:    "+",
		BottomLeft:  "+",
		BottomRight: "+",
		Horizontal:  "-",
		Vertical:    "|",
		Color:       Magenta,
	}

	DoubleBoxStyle = BoxStyle{
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
		Horizontal:  "═",
		Vertical:    "║",
		Color:       Cyan,
	}

	RoundBoxStyle = BoxStyle{
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
		Horizontal:  "─",
		Vertical:    "│",
		Color:       Green,
	}

	ThickBoxStyle = BoxStyle{
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
		Horizontal:  "━",
		Vertical:    "┃",
		Color:       LightBlue,
	}
)

// drawBox creates a box with the given style and text
func drawBox(text string, style BoxStyle, width int) string {
	if width <= 4 {
		width = len(text) + 4
	}

	// Create top border
	topBorder := style.Color + style.TopLeft + strings.Repeat(style.Horizontal, width-2) + style.TopRight + Reset

	// Create bottom border
	bottomBorder := style.Color + style.BottomLeft + strings.Repeat(style.Horizontal, width-2) + style.BottomRight + Reset

	// Create middle line with text (centered)
	innerWidth := width - 4 // Account for left and right borders + spaces
	paddingTotal := innerWidth - len(text)
	paddingLeft := paddingTotal / 2
	paddingRight := paddingTotal - paddingLeft

	middle := style.Color + style.Vertical + Reset +
		" " +
		strings.Repeat(" ", paddingLeft) + text + strings.Repeat(" ", paddingRight) +
		" " +
		style.Color + style.Vertical + Reset

	return topBorder + "\n" + middle + "\n" + bottomBorder
}

// SimpleBox creates a box with simple ASCII characters
func SimpleBox(text string) string {
	return drawBox(text, SimpleBoxStyle, 80)
}

// PrintSimpleBox prints a simple box
func PrintSimpleBox(text string) {
	fmt.Println(drawBox(text, SimpleBoxStyle, 80))
}

// DoubleBox creates a box with double-line characters
func DoubleBox(text string) string {
	return drawBox(text, DoubleBoxStyle, 80)
}

// PrintDoubleBox prints a double box
func PrintDoubleBox(text string) {
	fmt.Println(drawBox(text, DoubleBoxStyle, 80))
}

// RoundBox creates a box with rounded corners
func RoundBox(text string) string {
	return drawBox(text, RoundBoxStyle, 80)
}

// PrintRoundBox prints a rounded box
func PrintRoundBox(text string) {
	fmt.Println(drawBox(text, RoundBoxStyle, 80))
}

// ThickBox creates a box with thick characters
func ThickBox(text string) string {
	return drawBox(text, ThickBoxStyle, 80)
}

// PrintThickBox prints a thick box
func PrintThickBox(text string) {
	fmt.Println(drawBox(text, ThickBoxStyle, 80))
}

// CustomBox creates a box with custom style and width
func CustomBox(text string, style BoxStyle, width int) string {
	return drawBox(text, style, width)
}

// PrintCustomBox prints a custom box
func PrintCustomBox(text string, style BoxStyle, width int) {
	fmt.Println(drawBox(text, style, width))
}
