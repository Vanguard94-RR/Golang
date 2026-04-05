package main

// ANSI color codes for terminal output
// Format: \x1b[<code>m

// Reset and Formatting
const (
	Reset        = "\x1b[0m"
	Bold         = "\x1b[1m"
	Italics      = "\x1b[3m"
	Underline    = "\x1b[4m"
	Blink        = "\x1b[5m"
	Inverse      = "\x1b[7m"
	BoldOff      = "\x1b[22m"
	ItalicsOff   = "\x1b[23m"
	UnderlineOff = "\x1b[24m"
	BlinkOff     = "\x1b[25m"
	InverseOff   = "\x1b[27m"
)

// Foreground Colors
const (
	Black           = "\x1b[30m"
	Red             = "\x1b[31m"
	Green           = "\x1b[32m"
	Yellow          = "\x1b[33m"
	Blue            = "\x1b[34m"
	Magenta         = "\x1b[35m"
	Cyan            = "\x1b[36m"
	Gray            = "\x1b[37m"
	DarkGray        = "\x1b[30;1m"
	LightRed        = "\x1b[31;1m"
	LightGreen      = "\x1b[32;1m"
	LightYellow     = "\x1b[33;1m"
	LightBlue       = "\x1b[34;1m"
	LightMagenta    = "\x1b[35;1m"
	LightCyan       = "\x1b[36;1m"
	White           = "\x1b[37;1m"
	ResetForeground = "\x1b[39m"
)

// Background Colors
const (
	BlackBackground   = "\x1b[40m"
	RedBackground     = "\x1b[41m"
	GreenBackground   = "\x1b[42m"
	YellowBackground  = "\x1b[43m"
	BlueBackground    = "\x1b[44m"
	MagentaBackground = "\x1b[45m"
	CyanBackground    = "\x1b[46m"
	GrayBackground    = "\x1b[47m"
	ResetBackground   = "\x1b[49m"
)
