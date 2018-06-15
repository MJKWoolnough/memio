package memio

import (
	"io"
	"testing"
)

var (
	_ io.Reader   = &LimitedBuffer{}
	_ io.Writer   = &LimitedBuffer{}
	_ io.WriterTo = &LimitedBuffer{}
)

func TestLimitedBufferWrite(t *testing.T) {
	data := []byte("Beep")
	writer := LimitedBuffer(data)[:0:4]
	if n, err := writer.Write([]byte("J")); n != 1 {
		t.Errorf("expecting to write 1 byte, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "Jeep" {
		t.Errorf("expecting %q, got %q", "Jeep  ", string(data))
		return
	}
	if n, err := writer.Write([]byte("ohn")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "John" {
		t.Errorf("expecting %q, got %q", "John  ", string(data))
		return
	}
	if n, err := writer.Write([]byte("ny")); err != io.ErrShortBuffer {
		t.Errorf("expecting io.ErrShortBuffer, got: %s", err)
		return
	} else if n != 0 {
		t.Errorf("expecting to write 0 bytes, read %d", n)
	}
}
