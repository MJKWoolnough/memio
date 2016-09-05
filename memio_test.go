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
		t.Errorf("expected to be at position 2, got %d", pos)
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
		t.Errorf("expected to be at position 9, got %d", pos)
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
	if pos, err := reader.Seek(-6, 2); pos != 7 {
		t.Errorf("expected to be at position 7, got %d", pos)
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
		t.Errorf("expected to be at position 0, got %d", pos)
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
	if pos, err := writer.Seek(-4, 2); pos != 2 {
		t.Errorf("expected to be at position 0, got %d", pos)
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
		t.Errorf("expected to be at position 6, got %d", pos)
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
	if err != ErrClosed {
		t.Errorf("expecting close error")
		return
	}
	_, err = writer.Write([]byte("Beep"))
	if err != ErrClosed {
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
	if n, err := writer.Write([]byte("World")); err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n != 5 {
		t.Errorf("expecting to write 5 bytes, wrote %d", n)
	} else if len(data) != 10 {
		t.Errorf("expecting buf to have 5 bytes, has %d", n)
	} else if string(data) != "HelloWorld" {
		t.Errorf("expecting %q, got %q", "Hello", string(data))
	}
}

func TestReadFrom(t *testing.T) {
	data := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	toRead := make([]byte, 0, 2480)
	for i := 0; i < 40; i++ {
		toRead = append(toRead, data...)
	}
	reader := Open(toRead)
	var toWrite []byte
	writer := Create(&toWrite)
	if n, err := writer.ReadFrom(reader); err != nil {
		t.Errorf("got error: %q", err.Error())
	} else if n != 2480 {
		t.Errorf("expecting to write 2480 bytes, wrote %d", n)
	} else if string(toRead) != string(toWrite) {
		t.Errorf("expecting %q, got %q", string(toRead), string(toWrite))
	}

}

func TestTruncate(t *testing.T) {
	data := make([]byte, 100)
	for i := byte(0); i < 100; i++ {
		data[i] = i % 10
	}
	w := Create(&data)
	w.Truncate(75)
	if len(data) != 75 {
		t.Errorf("expecting length 75, got %d", len(data))
		return
	}
	w.Truncate(90)
	if len(data) != 90 {
		t.Errorf("expecting length 90, got %d", len(data))
		return
	}
	for i := byte(0); i < 75; i++ {
		if data[i] != i%10 {
			t.Errorf("at position %d, expecting value of %d, got %d", i, i%10, data[i])
			return
		}
	}
	for i := byte(75); i < 90; i++ {
		if data[i] != 0 {
			t.Errorf("at position %d, expecting value of 0, got %d", i, data[i])
			return
		}
	}
	w.Truncate(100)
	if len(data) != 100 {
		t.Errorf("expecting length 100, got %d", len(data))
		return
	}
	for i := byte(90); i < 100; i++ {
		if data[i] != 0 {
			t.Errorf("at position %d, expecting value of 0, got %d", i, data[i])
			return
		}
	}
}
