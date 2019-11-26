package nexus

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_Print(t *testing.T) {
	var err error
	print := Print(err)
	print("anything")

	fprint := Fprint(err)
	fprint(ioutil.Discard, "Hello", "nexus!")

	printf := Printf(err)
	printf("Hello, %s!", "nexus")

	if err != nil {
		t.Error(err)
	}
}

func Test_Print_fails(t *testing.T) {
	err := fmt.Errorf("failed")
	buf := bytes.NewBufferString("")
	Stdout = buf

	print := Print(err)
	print("x")

	printf := Printf(err)
	printf("%s", "x")

	if err == nil {
		t.Error(err)
	}
	got := buf.String()
	if got != "" {
		t.Error(got)
	}
}
