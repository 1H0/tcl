package tcl

import (
	"fmt"
	"strings"
)

// defaultIndent specified the indent that is applied anyway,
// no matter what is specified as a parameter.
// If you call the 'indent()' function with a indent of '0', the indent specified here is applyed anyway.
const defaultIndent int = 2

// indentSize specifies the size of a single indentation.
// In the 'indent()' function, this value is multiplied by the provided indentation provided as a parameter.
const indentSize int = 2

// newLines does print any number of newlines provided
// as a parameter.
func NewLines(count int) {
	fmt.Print(strings.Repeat("\n", count))
}

// indent is a internal function that is used to indent text.
// The default indentation is set in the 'defaultIndent' variable.
// How much is indented per indent step is specified in the 'indentSize' variable.
func indent(indent int) {
	indent = indent*indentSize + defaultIndent
	fmt.Print(strings.Repeat(" ", indent))
}
