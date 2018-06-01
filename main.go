package main

import (
	"fmt"
	"strings"
)

type position struct {
	x, y int
}

var (
	rowLength       = 5
	currentPosition = position{x: 0, y: 0}
)

func main() {
	fmt.Printf("Example: %s\nTest: %s", wordToButtons("CAT"), wordToButtons("LEGOTHEBATMANMOVIEWITHSUPERMANANDNINJAS"))
}

func wordToButtons(s string) string {
	var fullButtonPath string
	for _, runeFromString := range s {
		fullButtonPath += printButtonPath(letterPosition(int(runeFromString) - int('A')))
	}
	// Reset position
	currentPosition = position{x: 0, y: 0}
	return fullButtonPath
}

func letterPosition(letterIndex int) position {
	return position{x: letterIndex % rowLength, y: letterIndex / rowLength}
}

func printButtonPath(in position) string {
	var out string
	xDiff := in.x - currentPosition.x
	yDiff := in.y - currentPosition.y
	// The order of these matters, to avoid the whitespace present in the bottom right of the keypad
	if xDiff < 0 {
		out += strings.Repeat("L", xDiff*-1)
	}
	if yDiff < 0 {
		out += strings.Repeat("U", yDiff*-1)
	}
	if yDiff > 0 {
		out += strings.Repeat("D", yDiff)
	}
	if xDiff > 0 {
		out += strings.Repeat("R", xDiff)
	}
	currentPosition = in
	return out + "O"
}
