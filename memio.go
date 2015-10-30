// Package memio implements Read, Write, Seek, Close and other io methods for a byte slice.
package memio

import (
	"bytes"
	"errors"
	"io"
)

const (
	seekSet = iota
	seekCurr
	seekEnd
)

// ErrClosed is an error returned when trying to perform an operation after using Close().
var ErrClosed = errors.New("operation not permitted when closed")

// ReadMem holds a byte slice that can be used for many io interfaces
type ReadMem struct {
	*bytes.Reader
}

// Open uses a byte slice for reading. Implements io.Reader, io.Seeker,
// io.Closer, io.ReaderAt, io.ByteReader and io.WriterTo.
func Open(data []byte) ReadMem {
	return ReadMem{bytes.NewReader(data)}
}

// Close is a no-op func the lets ReadMem implement interfaces that require a
// Close method
func (ReadMem) Close() error {
	return nil
}

// WriteMem holds a pointer to a byte slice and allows numerous io interfaces
// to be used with it.
type WriteMem struct {
	data *[]byte
	pos  int
}

// Create uses a byte slice for writing. Implements io.Writer, io.Seeker,
// io.Closer, io.WriterAt, io.ByteWriter and io.ReaderFrom.
func Create(data *[]byte) *WriteMem {
	return &WriteMem{data, 0}
}

// Write is an implementation of the io.Writer interface
func (b *WriteMem) Write(p []byte) (int, error) {
	if b.data == nil {
		return 0, ErrClosed
	}
	b.setSize(b.pos + len(p))
	n := copy((*b.data)[b.pos:], p)
	b.pos += n
	return n, nil
}

// WriteAt is an implementation of the io.WriterAt interface
func (b *WriteMem) WriteAt(p []byte, off int64) (int, error) {
	if b.data == nil {
		return 0, ErrClosed
	}
	b.setSize(int(off) + len(p))
	return copy((*b.data)[off:], p), nil
}

// WriteByte is an implementation of the io.WriteByte interface
func (b *WriteMem) WriteByte(c byte) error {
	if b.data == nil {
		return ErrClosed
	}
	b.setSize(b.pos + 1)
	(*b.data)[b.pos] = c
	b.pos++
	return nil
}

// ReadFrom is an implementation of the io.ReaderFrom interface
func (b *WriteMem) ReadFrom(f io.Reader) (int64, error) {
	if b.data == nil {
		return 0, ErrClosed
	}
	var (
		c   int64
		n   int
		err error
	)
	buf := make([]byte, 1024)
	for {
		n, err = f.Read(buf)
		if n > 0 {
			c += int64(n)
			b.setSize(b.pos + n)
			copy((*b.data)[b.pos:], buf[:n])
			b.pos += n
		}
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}
	return c, err
}

// Seek is an implementation of the io.Seeker interface
func (b *WriteMem) Seek(offset int64, whence int) (int64, error) {
	if b.data == nil {
		return 0, ErrClosed
	}
	switch whence {
	case seekSet:
		b.pos = int(offset)
	case seekCurr:
		b.pos += int(offset)
	case seekEnd:
		b.pos = len(*b.data) + int(offset)
	}
	if b.pos < 0 {
		b.pos = 0
	}
	return int64(b.pos), nil
}

// Close is and implementation of the io.Closer interface
func (b *WriteMem) Close() error {
	b.data = nil
	return nil
}

func (b *WriteMem) setSize(end int) {
	if end > len(*b.data) {
		if end < cap(*b.data) {
			*b.data = (*b.data)[:end]
		} else {
			var newData []byte
			if len(*b.data) < 512 {
				newData = make([]byte, end, end<<1)
			} else {
				newData = make([]byte, end, end+(end>>2))
			}
			copy(newData, *b.data)
			*b.data = newData
		}
	}
}

// Truncate changes the length of the byte slice to the given amount
func (b *WriteMem) Truncate(s int64) error {
	if l := int64(len(*b.data)); l > s {
		copy((*b.data)[s:], make([]byte, l-s))
		*b.data = (*b.data)[:s]
	} else if l < s {
		b.setSize(int(s))
	}
	return nil
}

func (b *WriteMem) WriteString(s string) (int, error) {
	return b.Write([]byte(s))
}
