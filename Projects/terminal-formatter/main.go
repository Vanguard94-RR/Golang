package main

import "fmt"

func main() {
	fmt.Println("\n=== Terminal Formatter - All Styles ===\n")

	// Preset Styles - Quick and Easy
	fmt.Println("--- Quick Presets (easy to use) ---")
	PrintTitle("Welcome to Terminal Formatter")
	PrintSubTitle("This is a subtitle demo")
	PrintSection("Section Header")
	PrintSuccess("Operation completed successfully!")
	PrintWarning("This is a warning message")
	PrintError("An error has occurred!")
	PrintInfo("Here is some information")

	// Box Styles - All Types
	fmt.Println("\n--- Box Styles (with real borders) ---")
	PrintSimpleBox("Simple Box - ASCII Characters")
	fmt.Println()
	PrintRoundBox("Round Box - Rounded Corners")
	fmt.Println()
	PrintDoubleBox("Double Box - Double Line Characters")
	fmt.Println()
	PrintThickBox("Thick Box - Heavy Characters")

	fmt.Println("\n--- Custom Styling (full control) ---")

	// Example 1: Custom width
	New().WithWidth(60).WithColor(Blue).WithBold(true).Print("Custom Width 60")

	// Example 2: Different padding character
	New().WithPadChar("*").WithColor(Yellow).Print("Stars as Padding")

	// Example 3: Combination (Bold + Italic)
	New().
		WithColor(LightMagenta).
		WithBold(true).
		WithItalic(true).
		Print("Bold and Italic")

	// Example 4: Get string without printing
	fmt.Println("\n--- Returning Strings (for reuse) ---")
	formatter := New().WithColor(Cyan)
	formatted := formatter.Format("Reusable String")
	fmt.Println("String 1:", formatted)
	fmt.Println("String 2:", formatted) // Can reuse the same string!

	// Example 5: Using Simple color function
	fmt.Println("\n--- Simple Colors ---")
	PrintSimple("Red Text", Red)
	PrintSimple("Green Text", Green)
	PrintSimple("Blue Text", Blue)

	// Example 6: Custom box with specific style
	fmt.Println("\n--- Custom Box (specific width) ---")
	PrintCustomBox("Custom Box 100 width", DoubleBoxStyle, 100)

	fmt.Println("\n=== End Examples ===\n")
}
