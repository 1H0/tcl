package tcl

import (
	"fmt"
	"strings"

	c "github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// Title prints a title with backgorund and text in uppercase.
func Title(title string) {
	titleColor := c.New(c.BgBlue, c.FgBlack, c.Bold)

	indent(0)
	titleColor.Printf("  %s  ", strings.ToUpper(title))
	NewLines(2)

}

// TitleC is like Title() but is centered in the console.
func TitleC(title string) {
	titleColor := c.New(c.BgBlue, c.FgBlack, c.Bold)
	width, _, _ := terminal.GetSize(0)

	leftOffset := (width - len(title)) / 2

	indent(0)
	fmt.Print(strings.Repeat(" ", leftOffset))
	titleColor.Printf("  %s  ", strings.ToUpper(title))
	NewLines(2)
}

// SubTitle prints a subtitle that is bold and underlined.
func SubTitle(title string) {
	titleColor := c.New(c.FgBlue, c.Bold, c.Underline)

	indent(0)
	titleColor.Print(title)
	NewLines(1)
}

// FullTitle prints out a title that is as wide as your terminal is.
func FullTitle(title string) {
	width, _, _ := terminal.GetSize(0)
	titleColor := c.New(c.BgBlue, c.FgBlack, c.Bold)

	indent(0)
	titleColor.Printf("  %s", strings.ToUpper(title))
	titleColor.Print(strings.Repeat(" ", width-len(title)-6))
	NewLines(2)
}
