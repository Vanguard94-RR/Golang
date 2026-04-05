package main

import (
	"fmt"
	"strings"
)

// TextFormatter defines how to format text in the terminal
type TextFormatter struct {
	Width     int    // Total width (default 80)
	PadChar   string // Character for padding (default "-")
	TextColor string // ANSI color code
	Bold      bool   // Apply bold formatting
	Italic    bool   // Apply italic formatting
	Centered  bool   // Center the text (default true)
}

// New creates a default TextFormatter
func New() *TextFormatter {
	return &TextFormatter{
		Width:     80,
		PadChar:   "-",
		TextColor: "",
		Bold:      false,
		Italic:    false,
		Centered:  true,
	}
}

// WithWidth sets the formatter width
func (tf *TextFormatter) WithWidth(width int) *TextFormatter {
	tf.Width = width
	return tf
}

// WithColor sets the text color
func (tf *TextFormatter) WithColor(color string) *TextFormatter {
	tf.TextColor = color
	return tf
}

// WithBold enables bold formatting
func (tf *TextFormatter) WithBold(bold bool) *TextFormatter {
	tf.Bold = bold
	return tf
}

// WithItalic enables italic formatting
func (tf *TextFormatter) WithItalic(italic bool) *TextFormatter {
	tf.Italic = italic
	return tf
}

// WithPadChar sets the padding character
func (tf *TextFormatter) WithPadChar(char string) *TextFormatter {
	tf.PadChar = char
	return tf
}

// Format returns the formatted string WITHOUT printing
// This is key - it allows reuse, testing, and chaining
func (tf *TextFormatter) Format(text string) string {
	if !tf.Centered {
		return text
	}

	// Calculate padding
	leftDashes := (tf.Width - len(text)) / 2
	rightDashes := tf.Width - len(text) - leftDashes

	// Build left and right padding
	left := strings.Repeat(tf.PadChar, leftDashes)
	right := strings.Repeat(tf.PadChar, rightDashes)

	// Apply formatting
	result := left + text + right

	// Apply colors and styles
	if tf.Bold {
		result = Bold + result
	}
	if tf.Italic {
		result = Italics + result
	}
	if tf.TextColor != "" {
		result = tf.TextColor + result
	}
	result += Reset

	return result
}

// Print prints the formatted text
func (tf *TextFormatter) Print(text string) {
	fmt.Println(tf.Format(text))
}

// String implements the Stringer interface
func (tf *TextFormatter) String() string {
	return fmt.Sprintf("TextFormatter{Width:%d, Color:%s, Bold:%v, Italic:%v}",
		tf.Width, tf.TextColor, tf.Bold, tf.Italic)
}
