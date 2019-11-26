package nexus

import (
	"fmt"
	"io"
	"os"
)

var Stdout io.Writer = os.Stdout

func Print(err error) func(...interface{}) int {
	return func(args ...interface{}) int {
		if err != nil {
			return 0
		}
		var n int
		n, err = fmt.Fprint(Stdout, args...)
		return n
	}
}
