package memio

import (
	"io"
	"testing"
)

var (
	_ io.Reader     = new(ReadMem)
	_ io.Closer     = new(ReadMem)
	_ io.Seeker     = new(ReadMem)
	_ io.WriterTo   = new(ReadMem)
	_ io.ByteReader = new(ReadMem)
	_ io.ReaderAt   = new(ReadMem)

	_ io.Writer     = new(WriteMem)
	_ io.Closer     = new(WriteMem)
	_ io.Seeker     = new(WriteMem)
	_ io.ReaderFrom = new(WriteMem)
	_ io.ByteWriter = new(WriteMem)
	_ io.WriterAt   = new(WriteMem)
)

func TestReadWriteRead(t *testing.T) {
	data := []byte("Hello, World!")
	reader := OpenMem(&data)
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
	} else if pos, err := reader.Seek(2, 0); pos != 2 {
		t.Errorf("expected to be at position 2, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n, err = reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(toRead) != "llo, " {
		t.Errorf("expecting %q, got %q", "llo, ", string(toRead))
	} else if pos, err = reader.Seek(2, 1); pos != 9 {
		t.Errorf("expected to be at position 9, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n, err := reader.Read(toRead); n != 4 {
		t.Errorf("expecting to read 4 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(toRead) != "rld! " {
		t.Errorf("expecting %q, got %q", "rld! ", string(toRead))
	} else if pos, err = reader.Seek(-6, 2); pos != 7 {
		t.Errorf("expected to be at position 7, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n, err := reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(toRead) != "World" {
		t.Errorf("expecting %q, got %q", "World", string(toRead))
	} else if _, err = reader.Seek(1, 0); err != nil {
		t.Errorf("got error: %q", err.Error())
	}

	reader.Close()
}

func TestReadWriteWrite(t *testing.T) {
	data := []byte("Beep")
	writer := OpenMem(&data)

	if n, err := writer.Write([]byte("J")); n != 1 {
		t.Errorf("expecting to write 1 byte, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Jeep" {
		t.Errorf("expecting %q, got %q", "Jeep", string(data))
	} else if n, err = writer.Write([]byte("ohn")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "John" {
		t.Errorf("expecting %q, got %q", "John", string(data))
	} else if n, err = writer.Write([]byte("ny")); n != 2 {
		t.Errorf("expecting to write 2 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Johnny" {
		t.Errorf("expecting %q, got %q", "Johnny", string(data))
	} else if pos, err := writer.Seek(0, 0); pos != 0 {
		t.Errorf("expected to be at position 0, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n, err = writer.Write([]byte("Edmund")); n != 6 {
		t.Errorf("expecting to write 6 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Edmund" {
		t.Errorf("expecting %q, got %q", "Edmund", string(data))
	} else if pos, err = writer.Seek(-4, 2); pos != 2 {
		t.Errorf("expected to be at position 0, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n, err = writer.Write([]byte("war")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Edward" {
		t.Errorf("expecting %q, got %q", "Edward", string(data))
	} else if pos, err = writer.Seek(1, 1); pos != 6 {
		t.Errorf("expected to be at position 6, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n, err = writer.Write([]byte("o")); n != 1 {
		t.Errorf("expecting to write 1 bytes, wrote %d", n)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if string(data) != "Edwardo" {
		t.Errorf("expecting %q, got %q", "Edwardo", string(data))
	} else if err = writer.Close(); err != nil {
		t.Errorf("unexpected error: %s", err)
	} else if _, err = writer.Seek(0, 0); err != ErrClosed {
		t.Errorf("expecting close error")
	} else if _, err = writer.Write([]byte("Beep")); err != ErrClosed {
		t.Errorf("expecting close error")
	}
}
