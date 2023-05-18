package nexus

import (
	"fmt"
	"io"
	"os"
)

func ExampleStepper() {
	var (
		src  = "input.txt"
		dst  = "output.txt"
		in   *os.File
		out  *os.File
		err  error
		next = NewStepper(&err)
	)

	next.Step(func() {
		in, err = os.Open(src)
	})

	next.Stepf("create", func() {
		out, err = os.Create(dst)
	})

	next.Stepf("copy", func() {
		_, err = io.Copy(out, in)
	})

	fmt.Println(err)
	// output:
	// open input.txt: no such file or directory
}
