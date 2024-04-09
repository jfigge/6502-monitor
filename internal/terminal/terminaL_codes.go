/*
 * Copyright (C) 2024 by Jason Figge
 */

package terminal

// Screen control
const (
	ScreenCls = "\u001Bc\u001b[?25l"
	EOL       = "\u001B[K"

	CursorAt    = "\u001B[%d;%dH%s"
	CursorShow  = "\u001B[?25h"
	CursorHide  = "\u001B[?25l"
	CursorUp    = "\u001B[%dA"
	CursorDown  = "\u001B[%dB"
	CursorRight = "\u001B[%dC"
	CursorLeft  = "\u001B[%dD"
	CursorSave  = "\u001B7"
	CursorLoad  = "\u001B8"
)

// Colors
const (
	Black   = "\u001B[30m"
	Red     = "\u001B[31m"
	Green   = "\u001B[32m"
	Yellow  = "\u001B[33m"
	Blue    = "\u001B[34m"
	Magenta = "\u001B[35m"
	Cyan    = "\u001B[36m"
	White   = "\u001B[37m"

	Grey          = "\u001B[90m"
	BrightRed     = "\u001B[91m"
	BrightGreen   = "\u001B[92m"
	BrightYellow  = "\u001B[93m"
	BrightBlue    = "\u001B[94m"
	BrightMagenta = "\u001B[95m"
	BrightCyan    = "\u001B[96m"
	BrightWhite   = "\u001B[97m"
	Reset         = "\u001B[0m"
)

// Keycodes
const (
	KeyCTRLC  = 3
	KeyKey    = 0
	KeyUp     = 38
	KeyOption = 1
	KeyDown   = 40
	KeyLeft   = 37
	KeyRight  = 39
)

// Characters
const (
	Asc   = "↑"
	Desc  = "↓"
	Check = "✓"
	Cross = "x"
)
