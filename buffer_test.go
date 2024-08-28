package memio

import (
	"bytes"
	"io"
	"testing"
)

var (
	_ io.Reader   = &Buffer{}
	_ io.Writer   = &Buffer{}
	_ io.WriterTo = &Buffer{}
	_ io.ReaderAt = &Buffer{}
	_ io.WriterAt = &Buffer{}
)

func TestBufferRead(t *testing.T) {
	reader := Buffer("Hello, World!")
	toRead := make([]byte, 5)

	if n, err := reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(toRead) != "Hello" {
		t.Errorf("expecting %q, got %q", "Hello", string(toRead))
	} else if n, err = reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(toRead) != ", Wor" {
		t.Errorf("expecting %q, got %q", ", Wor", string(toRead))
	} else if n, err = reader.Read(toRead); n != 3 {
		t.Errorf("expecting to read 3 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(toRead) != "ld!or" {
		t.Errorf("expecting %q, got %q", "ld!or", string(toRead))
	} else if n, err = reader.Read(toRead); n != 0 {
		t.Errorf("expecting to read 0 bytes, read %d", n)
	} else if err != io.EOF {
		t.Errorf("expecting EOF")
	}
}

func TestBufferWrite(t *testing.T) {
	data := []byte("Beep  ")
	writer := Buffer(data)[:0]

	if n, err := writer.Write([]byte("J")); n != 1 {
		t.Errorf("expecting to write 1 byte, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Jeep  " {
		t.Errorf("expecting %q, got %q", "Jeep  ", string(data))
	} else if n, err = writer.Write([]byte("ohn")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "John  " {
		t.Errorf("expecting %q, got %q", "John  ", string(data))
	} else if n, err := writer.Write([]byte("ny")); n != 2 {
		t.Errorf("expecting to write 2 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Johnny" {
		t.Errorf("expecting %q, got %q", "Johnny", string(data))
	}
}

func TestBufferReadFrom(t *testing.T) {
	for n, test := range [...]struct {
		byteReader
		limit   int64
		initial Buffer
		result  []byte
	}{
		/*
			{
				3,
				23,
				make(Buffer, 10, 10),
				[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1},
			},
			{
				3,
				31,
				make(Buffer, 0, 0),
				[]byte{0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1},
			},
		*/
	} {
		if m, err := test.initial.ReadFrom(io.LimitReader(&test.byteReader, test.limit)); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else if m != test.limit {
			t.Errorf("test %d: expecting to read %d bytes, read %d", n+1, test.limit, m)
		} else if !bytes.Equal(test.result, test.initial) {
			t.Errorf("test %d: expecting %v, got %v", n+1, test.result, test.initial)
		}
	}
}
