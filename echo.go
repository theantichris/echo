package echo

import (
	"fmt"
	"io"
)

// Echo writes out a list of strings separated by spaces.
func Echo(args []string, w io.Writer) {
	var output, separator string

	for _, arg := range args[1:] {
		output += separator + arg
		separator = " "
	}

	fmt.Fprintln(w, output)
}
