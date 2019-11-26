package nexus

import (
	"fmt"
	"io"
	"os"
)

var Stdout io.Writer = os.Stdout

func Print(err error) func(...interface{}) int {
	fprint := Fprint(err)
	return func(args ...interface{}) int {
		return fprint(Stdout, args...)
	}
}

func Fprint(err error) func(io.Writer, ...interface{}) int {
	return func(w io.Writer, args ...interface{}) int {
		if err != nil {
			return 0
		}
		var n int
		n, err = fmt.Fprint(w, args...)
		return n
	}
}
