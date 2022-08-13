package nexus

import (
	"encoding"
	"fmt"
	"io"
)

// NewWriter returns a writer nexus. The referenced error is set by
// all methods if an error occurs.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

type Writer struct {
	w   io.Writer
	Err error
	// Total number of bytes written
	Written int64
}

// WriteBinary writes binary data to the underlying writer.  If given
// value implements encoding.BinaryMarshaler, the bytes from
// marshaling are written.
// Types []byte and byte are also recognized, panics on other data.
func (t *Writer) WriteBinary(m ...interface{}) {
	if t.Err != nil {
		return
	}
	for _, m := range m {
		switch m := m.(type) {

		case encoding.BinaryMarshaler:
			data, err := m.MarshalBinary()
			if err != nil {
				t.Err = err
				return
			}
			t.Write(data)

		case []byte:
			t.Write(m)

		case byte:
			t.Write([]byte{m})

		default:
			panic(fmt.Errorf("WriteBinary %T not allowed, convert to []byte", m))
		}
	}
}

// Write writes the bytes using the underlying writer. Does nothing if
// previous writes have failed.
func (t *Writer) Write(b []byte) (int, error) {
	if t.Err != nil {
		return 0, t.Err
	}
	var n int
	n, t.Err = t.w.Write(b)
	t.Written += int64(n)
	return n, t.Err
}

// Done returns total written bytes and error if any
func (t *Writer) Done() (int64, error) {
	return t.Written, t.Err
}
