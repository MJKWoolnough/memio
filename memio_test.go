package memio

import (
	"io"
	"testing"
)

var (
	_ io.Reader     = new(readMem)
	_ io.Closer     = new(readMem)
	_ io.Seeker     = new(readMem)
	_ io.WriterTo   = new(readMem)
	_ io.ByteReader = new(readMem)
	_ io.ReaderAt   = new(readMem)

	_ io.Writer     = new(writeMem)
	_ io.Closer     = new(writeMem)
	_ io.Seeker     = new(writeMem)
	_ io.ReaderFrom = new(writeMem)
	_ io.ByteWriter = new(writeMem)
	_ io.WriterAt   = new(writeMem)
)

func TestRead(t *testing.T) {
	data := []byte("Hello, World!")
	reader := Open(data)
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
	if pos, err := reader.Seek(2, 0); pos != 2 {
		t.Errorf("expected to be at postion 2, got %d", pos)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if n, err := reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(toRead) != "llo, " {
		t.Errorf("expecting %q, got %q", "llo, ", string(toRead))
		return
	}
	if pos, err := reader.Seek(2, 1); pos != 9 {
		t.Errorf("expected to be at postion 9, got %d", pos)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if n, err := reader.Read(toRead); n != 4 {
		t.Errorf("expecting to read 4 bytes, read %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(toRead) != "rld! " {
		t.Errorf("expecting %q, got %q", "rld! ", string(toRead))
		return
	}
	if pos, err := reader.Seek(6, 2); pos != 7 {
		t.Errorf("expected to be at postion 7, got %d", pos)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if n, err := reader.Read(toRead); n != 5 {
		t.Errorf("expecting to read 5 bytes, read %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(toRead) != "World" {
		t.Errorf("expecting %q, got %q", "World", string(toRead))
		return
	}
	if _, err := reader.Seek(1, 0); err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	}
	reader.Close()
	_, err := reader.Seek(1, 0)
	if _, ok := err.(*Closed); !ok {
		t.Errorf("expecting close error")
		return
	}
	_, err = reader.Read(toRead)
	if _, ok := err.(*Closed); !ok {
		t.Errorf("expecting close error")
		return
	}
}

func TestWrite(t *testing.T) {
	data := []byte("Beep")
	writer := Create(&data)
	if n, err := writer.Write([]byte("J")); n != 1 {
		t.Errorf("expecting to write 1 byte, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "Jeep" {
		t.Errorf("expecting %q, got %q", "Jeep", string(data))
		return
	}
	if n, err := writer.Write([]byte("ohn")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "John" {
		t.Errorf("expecting %q, got %q", "John", string(data))
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
	if pos, err := writer.Seek(0, 0); pos != 0 {
		t.Errorf("expected to be at postion 0, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if n, err := writer.Write([]byte("Edmund")); n != 6 {
		t.Errorf("expecting to write 6 bytes, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "Edmund" {
		t.Errorf("expecting %q, got %q", "Edmund", string(data))
		return
	}
	if pos, err := writer.Seek(4, 2); pos != 2 {
		t.Errorf("expected to be at postion 0, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if n, err := writer.Write([]byte("war")); n != 3 {
		t.Errorf("expecting to write 3 bytes, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "Edward" {
		t.Errorf("expecting %q, got %q", "Edward", string(data))
		return
	}
	if pos, err := writer.Seek(1, 1); pos != 6 {
		t.Errorf("expected to be at postion 6, got %d", pos)
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if n, err := writer.Write([]byte("o")); n != 1 {
		t.Errorf("expecting to write 1 bytes, wrote %d", n)
		return
	} else if err != nil {
		t.Errorf("got error: %q", err.Error())
		return
	} else if string(data) != "Edwardo" {
		t.Errorf("expecting %q, got %q", "Edwardo", string(data))
		return
	}
	writer.Close()
	_, err := writer.Seek(0, 0)
	if _, ok := err.(*Closed); !ok {
		t.Errorf("expecting close error")
		return
	}
	_, err = writer.Write([]byte("Beep"))
	if _, ok := err.(*Closed); !ok {
		t.Errorf("expecting close error")
		return
	}
}

func TestNewWrite(t *testing.T) {
	var data []byte
	writer := Create(&data)
	if n, err := writer.Write([]byte("Hello")); err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n != 5 {
		t.Errorf("expecting to write 5 bytes, wrote %d", n)
	} else if len(data) != 5 {
		t.Errorf("expecting buf to have 5 bytes, has %d", n)
	} else if string(data) != "Hello" {
		t.Errorf("expecting %q, got %q", "Hello", string(data))
	}
}
