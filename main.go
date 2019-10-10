// Echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var output, separator string

	for i := 1; i < len(os.Args); i++ {
		output += separator + os.Args[i]
		separator = " "
	}

	fmt.Println(output)
}
