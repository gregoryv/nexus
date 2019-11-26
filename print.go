package nexus

import (
	"fmt"
	"io"
	"os"
)

// Stdout is the default destination of print funcs.
var Stdout io.Writer = os.Stdout

// Print returns a fmt.Print wrapper using err for errors.
func Print(err *error) func(...interface{}) int {
	fprint := Fprint(err)
	return func(args ...interface{}) int {
		return fprint(Stdout, args...)
	}
}

// Fprint returns a fmt.Fprint wrapper using err for errors.
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

// Printf returns a fmt.Printf wrapper using err for errors.
func Printf(err *error) func(string, ...interface{}) int {
	fprintf := Fprintf(err)
	return func(format string, args ...interface{}) int {
		return fprintf(Stdout, format, args...)
	}
}

// Fprintf returns a fmt.Fprintf wrapper using err for errors.
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

// Println returns a fmt.Println wrapper using err for errors.
func Println(err *error) func(...interface{}) int {
	fprintln := Fprintln(err)
	return func(args ...interface{}) int {
		return fprintln(Stdout, args...)
	}
}

// Fprintln returns a fmt.Println wrapper using err for errors.
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
