package echo

import (
	"fmt"
	"io"
)

// Echo writes out a list of strings separated by spaces.
func Echo(args []string, w io.Writer) {
	var output, separator string

	for i := 1; i < len(args); i++ {
		output += separator + args[i]
		separator = " "
	}

	fmt.Fprintln(w, output)
}
