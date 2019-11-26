package nexus

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_Print(t *testing.T) {
	var err *error
	print := Print(err)
	print("anything")

	fprint := Fprint(err)
	fprint(ioutil.Discard, "Hello", "nexus!")

	printf := Printf(err)
	printf("Hello, %s!", "nexus")

	fprintf := Fprintf(err)
	fprintf(ioutil.Discard, "Hello, %s!", "nexus")

	if err != nil {
		t.Error(err)
	}
}

func Test_Print_fails(t *testing.T) {
	buf := bytes.NewBufferString("")
	Stdout = buf

	var err *error
	print := Print(err)
	print("ok")

	e := fmt.Errorf("failed")
	err = &e
	print("x")

	printf := Printf(err)
	printf("%s", "x")

	got := buf.String()
	if got != "ok" {
		t.Error(got)
	}
}
