# Terminal Formatter

A lightweight, reusable Go library for formatting and colorizing terminal output.

## Features

- ✅ **Chainable API**: Fluent interface for configuration
- ✅ **Preset Styles**: Quick functions for common use cases (Title, Error, Success, etc)
- ✅ **Customizable**: Configure width, colors, padding, bold, italic
- ✅ **Reusable**: Returns strings, not just prints
- ✅ **ANSI Colors**: Full support for terminal colors and formatting
- ✅ **Type-Safe**: No magic strings, uses struct fields

## Installation

```bash
go get github.com/yourusername/terminal-formatter
```

## Quick Start

### Using Presets (Easiest)

```go
package main

import "terminal-formatter"

func main() {
    formatter.PrintTitle("My Application")
    formatter.PrintSuccess("Operation completed!")
    formatter.PrintError("An error occurred!")
}
```

### Using Builder Pattern (Full Control)

```go
formatter := formatter.New().
    WithColor(formatter.Blue).
    WithBold(true).
    WithWidth(100)

formatter.Print("Custom Title")
```

### Using Format (Returns String)

```go
formatted := formatter.New().
    WithColor(formatter.Red).
    Format("Error Message")

// Now you can reuse it, log it, etc
fmt.Println(formatted)
```

## Available Presets

| Function | Color | Style | Padding |
|----------|-------|-------|---------|
| `Title()` | Green | Bold | `-` |
| `SubTitle()` | Cyan | Italic | `-` |
| `Section()` | Blue | Bold | `=` |
| `Success()` | Light Green | Bold | `-` |
| `Error()` | Light Red | Bold | `-` |
| `Warning()` | Light Yellow | Bold | `-` |
| `Info()` | Light Blue | Normal | `-` |

## Box Styles

Real boxes with borders! 4 different styles:

```go
SimpleBox("Message")    // +--------+ (ASCII)
RoundBox("Message")     // ╭────────╮ (Rounded)
DoubleBox("Message")    // ╔════════╗ (Double)
ThickBox("Message")     // ┏━━━━━━━━┓ (Thick)
```

### Output Examples

**Simple Box:**
```
+------------------+
|   Simple Box     |
+------------------+
```

**Round Box:**
```
╭──────────────────╮
│   Round Box      │
╰──────────────────╯
```

**Double Box:**
```
╔══════════════════╗
║   Double Box     ║
╚══════════════════╝
```

**Thick Box:**
```
┏━━━━━━━━━━━━━━━━━━┓
┃   Thick Box      ┃
┗━━━━━━━━━━━━━━━━━━┛
```

## Builder Methods

Every style can be customized:

```go
New().
    WithWidth(100).           // Set width (default 80)
    WithColor(Green).         // Set ANSI color
    WithBold(true).           // Enable bold
    WithItalic(true).         // Enable italic
    WithPadChar("=").         // Padding character (default "-")
    Print("My Text")          // Print formatted text
```

## Examples

### Example 1: Title with Custom Width

```go
formatter.New().
    WithWidth(120).
    WithColor(formatter.Blue).
    WithBold(true).
    Print("Wide Title")
```

**Output:**
```
---------------------------Wide Title---------------------------
```

### Example 2: Section Header

```go
formatter.Section("Database Operations")
```

### Example 3: Colored Text without Borders

```go
formatter.Simple("Just colored text", formatter.LightGreen)
```

### Example 4: Reusing Formatted Strings

```go
errorMsg := formatter.Error("Invalid input")
log.Error(errorMsg)
saveToFile(errorMsg)
```

### Example 5: Box Styles

```go
formatter.PrintSimpleBox("ASCII Box")
formatter.PrintRoundBox("Rounded Corners")
formatter.PrintDoubleBox("Double Lines")
formatter.PrintThickBox("Thick Borders")
```

## Available Colors

### Foreground Colors
- Basic: `Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `Gray`, `White`
- Light: `LightRed`, `LightGreen`, `LightYellow`, `LightBlue`, `LightMagenta`, `LightCyan`, `DarkGray`

### Background Colors
- `BlackBackground`, `RedBackground`, `GreenBackground`, `YellowBackground`, `BlueBackground`, `MagentaBackground`, `CyanBackground`, `GrayBackground`

### Formatting
- `Bold`, `BoldOff`
- `Italics`, `ItalicsOff`
- `Underline`, `UnderlineOff`
- `Reset` (clears all formatting)

## API Overview

### Main Type

```go
type TextFormatter struct {
    Width      int    // Total width (default 80)
    PadChar    string // Padding character (default "-")
    TextColor  string // ANSI color code
    Bold       bool   // Apply bold
    Italic     bool   // Apply italic
    Centered   bool   // Center text (default true)
}
```

### Core Methods

```go
New() *TextFormatter                    // Create default formatter
(tf *TextFormatter) Format(text) string // Get formatted string
(tf *TextFormatter) Print(text)         // Print formatted text

// Builder methods (all return *TextFormatter)
WithWidth(int)
WithColor(string)
WithBold(bool)
WithItalic(bool)
WithPadChar(string)
```

### Preset Functions

```go
Title(text) string
PrintTitle(text)
SubTitle(text) string
PrintSubTitle(text)
Section(text) string
PrintSection(text)
Success(text) string
PrintSuccess(text)
Error(text) string
PrintError(text)
Warning(text) string
PrintWarning(text)
Info(text) string
PrintInfo(text)
Simple(text, color) string
PrintSimple(text, color)

// Box functions (real borders!)
SimpleBox(text) string
PrintSimpleBox(text)
RoundBox(text) string
PrintRoundBox(text)
DoubleBox(text) string
PrintDoubleBox(text)
ThickBox(text) string
PrintThickBox(text)
CustomBox(text, style, width) string
PrintCustomBox(text, style, width)
```

## Design Patterns Used

### 1. Builder Pattern
Chainable methods return `*TextFormatter` for fluent API:
```go
New().WithColor(Red).WithBold(true).Print("text")
```

### 2. Functional Options
Preset functions encapsulate common configurations:
```go
Title("text")  // Pre-configured Green + Bold
```

### 3. Separation of Concerns
- `colors.go`: ANSI constants
- `formatter.go`: Core logic
- `styles.go`: Preset functions
- `main.go`: Examples

## Testing

Run examples:
```bash
cd terminal-formatter
go run .
```

## Future Enhancements

- [ ] Unit tests
- [ ] Alignment options (left, center, right)
- [ ] Multi-line box support
- [ ] Custom color definitions
- [ ] Background color support for text
- [ ] Box composition (nested boxes)

## License

MIT

## Contributing

Contributions welcome! Please submit PRs or issues.
