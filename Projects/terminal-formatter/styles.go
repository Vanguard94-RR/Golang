package main

import "fmt"

// Preset styles for common use cases
// These are "factory" functions that return a configured TextFormatter

// Title returns a formatter for main titles
func Title(text string) string {
	return New().
		WithWidth(80).
		WithColor(Green).
		WithBold(true).
		Format(text)
}

// PrintTitle prints a formatted title
func PrintTitle(text string) {
	New().
		WithWidth(80).
		WithColor(Green).
		WithBold(true).
		Print(text)
}

// SubTitle returns a formatter for subtitles
func SubTitle(text string) string {
	return New().
		WithWidth(80).
		WithColor(Cyan).
		WithItalic(true).
		Format(text)
}

// PrintSubTitle prints a formatted subtitle
func PrintSubTitle(text string) {
	New().
		WithWidth(80).
		WithColor(Cyan).
		WithItalic(true).
		Print(text)
}

// Section returns a formatted section header
func Section(text string) string {
	return New().
		WithWidth(80).
		WithColor(Blue).
		WithBold(true).
		WithPadChar("=").
		Format(text)
}

// PrintSection prints a formatted section header
func PrintSection(text string) {
	New().
		WithWidth(80).
		WithColor(Blue).
		WithBold(true).
		WithPadChar("=").
		Print(text)
}

// Error returns a formatter for error messages
func Error(text string) string {
	return New().
		WithWidth(80).
		WithColor(LightRed).
		WithBold(true).
		Format(text)
}

// PrintError prints a formatted error message
func PrintError(text string) {
	New().
		WithWidth(80).
		WithColor(LightRed).
		WithBold(true).
		Print(text)
}

// Success returns a formatter for success messages
func Success(text string) string {
	return New().
		WithWidth(80).
		WithColor(LightGreen).
		WithBold(true).
		Format(text)
}

// PrintSuccess prints a formatted success message
func PrintSuccess(text string) {
	New().
		WithWidth(80).
		WithColor(LightGreen).
		WithBold(true).
		Print(text)
}

// Warning returns a formatter for warning messages
func Warning(text string) string {
	return New().
		WithWidth(80).
		WithColor(LightYellow).
		WithBold(true).
		Format(text)
}

// PrintWarning prints a formatted warning message
func PrintWarning(text string) {
	New().
		WithWidth(80).
		WithColor(LightYellow).
		WithBold(true).
		Print(text)
}

// Info returns a formatter for information messages
func Info(text string) string {
	return New().
		WithWidth(80).
		WithColor(LightBlue).
		Format(text)
}

// PrintInfo prints a formatted info message
func PrintInfo(text string) {
	New().
		WithWidth(80).
		WithColor(LightBlue).
		Print(text)
}

// Box returns a formatted box around text (using = characters)
func Box(text string) string {
	return New().
		WithWidth(80).
		WithColor(Magenta).
		WithPadChar("=").
		Format(text)
}

// PrintBox prints a formatted box
func PrintBox(text string) {
	New().
		WithWidth(80).
		WithColor(Magenta).
		WithPadChar("=").
		Print(text)
}

// Simple returns minimal formatting (just color)
func Simple(text, color string) string {
	return color + text + Reset
}

// PrintSimple prints colored text
func PrintSimple(text, color string) {
	fmt.Print(color + text + Reset + "\n")
}
