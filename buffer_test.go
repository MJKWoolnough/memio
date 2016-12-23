package memio

import (
	"io"
	"testing"
)

var (
	_ io.Reader   = &Buffer{}
	_ io.Writer   = &Buffer{}
	_ io.WriterTo = &Buffer{}
)

func TestBufferRead(t *testing.T) {
	reader := Buffer("Hello, World!")
	toRead := make([]byte, 5)
	if n, err := reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(toRead) != "Hello" {
		t.Errorf("expecting %q, got %q", "Hello", string(toRead))
		return
	}
	if n, err := reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(toRead) != ", Wor" {
		t.Errorf("expecting %q, got %q", ", Wor", string(toRead))
		return
	}
	if n, err := reader.Read(toRead); n != 3 {
		t.Errorf("expecting to read 3 bytes, read %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(toRead) != "ld!or" {
		t.Errorf("expecting %q, got %q", "ld!or", string(toRead))
		return
	}
	if n, err := reader.Read(toRead); n != 0 {
		t.Errorf("expecting to read 0 bytes, read %d", n)
		return
	} else if err != io.EOF {
		t.Errorf("expecting EOF")
	}
}

func TestBufferWrite(t *testing.T) {
	data := []byte("Beep  ")
	writer := Buffer(data)[:0]
	if n, err := writer.Write([]byte("J")); n != 1 {
		t.Errorf("expecting to write 1 byte, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "Jeep  " {
		t.Errorf("expecting %q, got %q", "Jeep  ", string(data))
		return
	}
	if n, err := writer.Write([]byte("ohn")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "John  " {
		t.Errorf("expecting %q, got %q", "John  ", string(data))
		return
	}
	if n, err := writer.Write([]byte("ny")); n != 2 {
		t.Errorf("expecting to write 2 bytes, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "Johnny" {
		t.Errorf("expecting %q, got %q", "Johnny", string(data))
		return
	}
}
