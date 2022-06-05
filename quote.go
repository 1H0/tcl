package tcl

import (
	"github.com/fatih/color"
)

// Quote let's you print a simple quote to the console.
func Quote(text string) {
	quoteColor := color.New(color.FgHiWhite)

	indent(0)
	quoteColor.Printf("│ %s  ", text)
	NewLines(2)
}
