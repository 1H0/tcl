package tcl

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Table(headings []string, data [][]string) {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for i, heading := range headings {

		if i == 0 {
			fmt.Fprintf(w, "   %s\t│", heading)
		} else if i == len(headings)-1 {
			fmt.Fprintf(w, " %s\t\n", heading)
		} else {
			fmt.Fprintf(w, " %s\t│", heading)
		}

	}

	for _, row := range data {
		for i, attr := range row {

			if i == 0 {
				fmt.Fprintf(w, "   %s\t│", attr)
			} else if i == len(row)-1 {
				fmt.Fprintf(w, " %s\t\n", attr)
			} else {
				fmt.Fprintf(w, " %s\t│", attr)
			}

		}
	}

	w.Flush()
	NewLines(1)

}
