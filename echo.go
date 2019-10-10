package echo

import (
	"fmt"
	"io"
	"strings"
)

// Echo writes out a list of strings separated by spaces.
func Echo(args []string, w io.Writer) {
	output := strings.Join(args[1:], " ")

	fmt.Fprintln(w, output)
}
