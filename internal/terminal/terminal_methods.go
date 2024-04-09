/*
 * Copyright (C) 2024 by Jason Figge
 */

package terminal

import (
	"fmt"
)

func PrintAtf(row int, col int, message string, args ...any) {
	hdr := fmt.Sprintf(CursorAt, row, col, message)
	_, _ = fmt.Printf(hdr, args...)
}

func Cls() {
	PrintAtf(0, 0, ScreenCls)
}

func HideCursor() {
	fmt.Print(CursorHide)
}

func ShowCursor() {
	fmt.Print(CursorShow)
}
