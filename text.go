package tcl

import (
	"github.com/fatih/color"
)

// Text simply prints the provided text to the console, while respecting the default indent set.
func Text(text string) {

	textColor := color.New(color.FgWhite)

	indent(0)
	textColor.Println(text)
	NewLines(1)
}
