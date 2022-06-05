package tcl

import (
	"fmt"

	"github.com/fatih/color"
)

// KeyValue lets you provide a 'key' and a 'value' as strings, which then will be printed as a pair to the console.
// Because the key-value pair is presetnted as a unordered list item, you can specify a indentation via the 'indentation' parameter.
func KeyValue(key string, value string, indentation int) {

	keyColor := color.New(color.FgCyan, color.Bold)
	valueColor := color.New(color.FgWhite)

	indent(indentation)
	fmt.Print("• ")
	keyColor.Print(key)
	fmt.Print(": ")
	valueColor.Print(value)
	NewLines(1)

}
