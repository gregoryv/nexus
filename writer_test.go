package nexus

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestWriter(t *testing.T) {
	w := NewWriter(ioutil.Discard)

	data := []byte{1, 2, 3}
	w.Write(data)
	w.WriteBinary(byte(1), data, binary(data))

	exp := int64(10)
	if got, _ := w.Done(); got != exp {
		t.Errorf("got %v, expected %v", got, exp)
	}

	w = NewWriter(ioutil.Discard)
	w.WriteBinary(binary(data), binary(data))
	if got, _ := w.Done(); got != 6 {
		t.Error("expected 6 written bytes")
	}

	w.WriteBinary(brokenBin("x"), binary(data))
	w.WriteBinary(data)
	w.Write(data)
	if _, err := w.Done(); err == nil {
		t.Error("expect error")
	}

}

type binary []byte

func (b binary) MarshalBinary() ([]byte, error) {
	return b, nil
}

type brokenBin string

func (b brokenBin) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("%s", b)
}
