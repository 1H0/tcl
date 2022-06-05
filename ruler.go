package tcl

import (
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// Ruler lets you print a horizontal ruler to the console consisting of the provided caracter(s).
func Ruler(c string) {

	rulerColor := color.New(color.Faint)
	width, _, _ := terminal.GetSize(0)

	if c == "" {
		c = "-"
	}

	rulerColor.Print(strings.Repeat(c, width/len(c)))
	NewLines(2)

}
