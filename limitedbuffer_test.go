package memio

import (
	"bytes"
	"io"
	"testing"
)

var (
	_ io.Reader     = &LimitedBuffer{}
	_ io.Writer     = &LimitedBuffer{}
	_ io.WriterTo   = &LimitedBuffer{}
	_ io.ReaderFrom = &LimitedBuffer{}
	_ io.ReaderAt   = &LimitedBuffer{}
	_ io.WriterAt   = &LimitedBuffer{}
)

func TestLimitedBufferWrite(t *testing.T) {
	data := []byte("Beep")
	writer := LimitedBuffer(data)[:0:4]

	if n, err := writer.Write([]byte("J")); n != 1 {
		t.Errorf("expecting to write 1 byte, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Jeep" {
		t.Errorf("expecting %q, got %q", "Jeep  ", string(data))
	} else if n, err = writer.Write([]byte("ohn")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "John" {
		t.Errorf("expecting %q, got %q", "John  ", string(data))
	} else if n, err = writer.Write([]byte("ny")); err != io.ErrShortWrite {
		t.Errorf("expecting io.ErrShortWrite, got: %s", err)
	} else if n != 0 {
		t.Errorf("expecting to write 0 bytes, wrote %d", n)
	}
}

type byteReader byte

func (b byteReader) Read(p []byte) (int, error) {
	if int(b) < len(p) {
		p = p[:b]
	}
	for i := byte(0); i < byte(len(p)); i++ {
		p[i] = i
	}

	return len(p), nil
}

func TestLimitedBufferReadFrom(t *testing.T) {
	l := make(LimitedBuffer, 0, 30)

	n, err := l.ReadFrom(byteReader(1))
	if n != 30 {
		t.Errorf("expecting to read 30 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("go unexpected error: %s", err)
	} else if !bytes.Equal(l, bytes.Repeat([]byte{0}, 30)) {
		t.Errorf("expecting 30 0's, got %v", l)
	}

	l = l[:0]

	if n, err = l.ReadFrom(byteReader(2)); n != 30 {
		t.Errorf("expecting to read 30 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("go unexpected error: %s", err)
	} else if !bytes.Equal(l, bytes.Repeat([]byte{0, 1}, 15)) {
		t.Errorf("expecting 15 [0, 1]'s, got %v", l)
	}

	l = l[:0]

	if n, err = l.ReadFrom(byteReader(3)); n != 30 {
		t.Errorf("expecting to read 30 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("go unexpected error: %s", err)
	} else if !bytes.Equal(l, bytes.Repeat([]byte{0, 1, 2}, 10)) {
		t.Errorf("expecting 10 [0, 1, 2]'s, got %v", l)
	}

	l = l[:0]

	if n, err = l.ReadFrom(io.LimitReader(byteReader(4), 20)); n != 20 {
		t.Errorf("expecting to read 20 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("go unexpected error: %s", err)
	} else if !bytes.Equal(l, bytes.Repeat([]byte{0, 1, 2, 3}, 5)) {
		t.Errorf("expecting 5 [0, 1, 2, 3]'s, got %v", l)
	}
}
