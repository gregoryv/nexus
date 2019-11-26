package nexus

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Print(t *testing.T) {
	var err error
	print := Print(err)
	print("anything")
	if err != nil {
		t.Error(err)
	}
}

func Test_Print_fails(t *testing.T) {
	err := fmt.Errorf("failed")
	print := Print(err)
	buf := bytes.NewBufferString("")
	Stdout = buf
	print("nothing")
	if err == nil {
		t.Error(err)
	}
	got := buf.String()
	if got != "" {
		t.Error(got)
	}
}
