package tcl

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// Header prints out a provided text as a centered message, which can be used to display program names.
// An example would be to display the program name, along with the current version as shown below
//
//	tcl.Header("My Program @ v0.1.0")
func Header(text string) {

	copyColor := color.New(color.Faint, color.Bold)
	width, _, _ := terminal.GetSize(0)

	leftOffset := (width - len(text)) / 2

	NewLines(1)
	fmt.Print(strings.Repeat(" ", leftOffset))
	copyColor.Print(text)
	NewLines(2)

}
