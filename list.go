package tcl

import (
	"fmt"

	"github.com/fatih/color"
)

// ListItemU takes a string and displays it as a unordered list item.
// With the indentation parameter you can specify if and how deep you want to indent the entry, which might be useful is you want to nest entries.
func ListItemU(text string, indentation int) {

	itemColor := color.New(color.FgWhite)

	indent(indentation)
	fmt.Print("• ")
	itemColor.Print(text)
	NewLines(1)

}

// ListItemO takes a string and displays it as a ordered list item.
// With the indentation parameter you can specify if and how deep you want to indent the entry, which might be useful is you want to nest entries.
// With the number parameter, you can specify the number that is displayed before the entry.
func ListItemO(text string, indentation int, number int) {

	numberColor := color.New(color.Faint)
	itemColor := color.New(color.FgWhite)

	indent(indentation)
	numberColor.Printf("%d. ", number)
	itemColor.Print(text)
	NewLines(1)

}

// ListO creates a ordered list from a provided array of strings.
func ListO(list []string, indentation int) {
	for i, item := range list {
		ListItemO(item, indentation, i+1)
	}
}

// ListU creates a unordered list from a provided array of strings.
func ListU(list []string, indentation int) {
	for _, item := range list {
		ListItemU(item, indentation)
	}
}
