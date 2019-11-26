package nexus

import (
	"fmt"
	"io"
	"os"
)

var Stdout io.Writer = os.Stdout

func Print(err *error) func(...interface{}) int {
	fprint := Fprint(err)
	return func(args ...interface{}) int {
		return fprint(Stdout, args...)
	}
}

func Fprint(err *error) func(io.Writer, ...interface{}) int {
	return func(w io.Writer, args ...interface{}) int {
		if err != nil {
			return 0
		}
		n, e := fmt.Fprint(w, args...)
		err = &e
		return n
	}
}

func Printf(err *error) func(string, ...interface{}) int {
	fprintf := Fprintf(err)
	return func(format string, args ...interface{}) int {
		return fprintf(Stdout, format, args...)
	}
}

func Fprintf(err *error) func(io.Writer, string, ...interface{}) int {
	return func(w io.Writer, format string, args ...interface{}) int {
		if err != nil {
			return 0
		}
		n, e := fmt.Fprintf(w, format, args...)
		err = &e
		return n
	}
}

func Println(err *error) func(...interface{}) int {
	fprintln := Fprintln(err)
	return func(args ...interface{}) int {
		return fprintln(Stdout, args...)
	}
}

func Fprintln(err *error) func(io.Writer, ...interface{}) int {
	return func(w io.Writer, args ...interface{}) int {
		if err != nil {
			return 0
		}
		n, e := fmt.Fprintln(w, args...)
		err = &e
		return n
	}
}
