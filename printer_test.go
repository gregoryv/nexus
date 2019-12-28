package nexus

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPrinter(t *testing.T) {
	buf := bytes.NewBufferString("")
	w, err := NewPrinter(buf)
	if err == nil {
		t.Fatal(err)
	}
	w.Printf("ok %s\n", "printf")
	w.Println("ok print")
	w.Write([]byte("ok Write\n"))

	w.err = fmt.Errorf("failed")
	w.Printf("%s should not print this", "printf")
	w.Print("print should not print this")
	w.Println("print should not print this")
	w.Write([]byte("Write should not write this"))
	if strings.Index(buf.String(), " not ") > -1 {
		t.Error(buf.String())
	}
}

func ExampleNewPrinter() {
	w, err := NewPrinter(os.Stdout)
	w.Print("hello")
	*err = fmt.Errorf("stop subsequent calls")
	w.Printf("cruel %s", "world")
	// output:
	// hello
}
