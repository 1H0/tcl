package tcl

import (
	"strings"

	c "github.com/fatih/color"
)

// Message let's you print a title along with a message to the console.
// With the color parameter, you can specify the background color of the title.
// This is useful if you want to underline a result of a action that was taken,
// for example a green message after a successful action.
// The possible and allowed colors are:
//   - black
//   - blue
//   - cyan
//   - green
//   - magenta
//   - red
//   - white
//   - yellow
// If you don't specify a valid color, a white message will be printed.
func Message(title string, message string, color string) {
	titleColor := c.New(c.FgBlack)
	messageColor := c.New(c.Italic)

	switch strings.ToLower(color) {

	case "black":
		titleColor.Add(c.BgBlack, c.FgWhite)
	case "blue":
		titleColor.Add(c.BgBlue)
	case "cyan":
		titleColor.Add(c.BgCyan)
	case "green":
		titleColor.Add(c.BgGreen)
	case "magenta":
		titleColor.Add(c.BgMagenta)
	case "red":
		titleColor.Add(c.BgRed)
	case "white":
		titleColor.Add(c.BgWhite)
	case "yellow":
		titleColor.Add(c.BgYellow)
	default:
		titleColor.Add(c.BgWhite)
	}

	NewLines(1)
	indent(0)
	titleColor.Printf(" %s ", strings.ToUpper(title))
	messageColor.Printf(" > %s\n", message)
	NewLines(1)

}
