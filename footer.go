package tcl

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// Footer prints a simple dashed line.
// It can be used at the end of your output
// to separate it from the following one.
func Footer() {

	footerColor := color.New(color.Faint)

	width, _, _ := terminal.GetSize(0)

	NewLines(1)
	footerColor.Print(strings.Repeat(" — ", width/3))
	NewLines(2)

}

// FooterC does alsmost the same as the Footer() function, but
// adds a copyright notice with the provided name and the
// current year in the center of the dashed line.
func FooterC(name string) {

	footerColor := color.New(color.Faint)
	copyColor := color.New(color.Faint, color.Bold)
	width, _, _ := terminal.GetSize(0)

	year := time.Now().Year()
	copyString := fmt.Sprintf(" [ © %d %s ] ", year, name)

	leftRight := (width - len(copyString)) / 2

	NewLines(1)
	footerColor.Print(strings.Repeat(" — ", leftRight/3))
	copyColor.Print(copyString)
	footerColor.Print(strings.Repeat(" — ", leftRight/3))
	NewLines(2)

}
