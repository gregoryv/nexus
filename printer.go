package nexus

import (
	"fmt"
	"io"
)

// NewPrinter returns a writer nexus. The referenced error is set by
// all methods if an error occurs.
func NewPrinter(w io.Writer) (*Printer, *error) {
	t := &Printer{w: w}
	return t, &t.err
}

type Printer struct {
	w   io.Writer
	err error
}

// Print prints arguments using the underlying writer. Does nothing if
// Printer has failed.
func (t *Printer) Print(args ...interface{}) {
	if t.err != nil {
		return
	}
	_, t.err = fmt.Fprint(t.w, args...)
}

// Printf prints a formated string using the underlying writer. Does
// nothing if Printer has failed.
func (t *Printer) Printf(format string, args ...interface{}) {
	if t.err != nil {
		return
	}
	_, t.err = fmt.Fprintf(t.w, format, args...)
}

// Println prints arguments using the underlying writer. Does nothing if
// Printer has failed.
func (t *Printer) Println(args ...interface{}) {
	if t.err != nil {
		return
	}
	_, t.err = fmt.Fprintln(t.w, args...)
}

// Write writes the bytes using the underlying writer. Does nothing if
// Printer has failed.
func (t *Printer) Write(b []byte) (int, error) {
	if t.err != nil {
		return 0, t.err
	}
	return t.w.Write(b)
}
