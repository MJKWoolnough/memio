package memio

import (
	"errors"
	"io"
)

// ReadWriteMem is a combination of both the ReadMem and WriteMem types,
// allowing both all reads and writes to the same underlying byte slice.
type ReadWriteMem struct {
	WriteMem
}

// OpenMem uses a byte slice for reading and writing. Implements io.Reader,
// io.Writer, io.Seeker, io.ReaderAt, io.ByteReader, io.WriterTo, io.WriterAt,
// io.ByteWriter and io.ReaderFrom.
func OpenMem(data *[]byte) *ReadWriteMem {
	return &ReadWriteMem{WriteMem{data, 0}}
}

// Peek reads the next n bytes without advancing the position
func (b *ReadWriteMem) Peek(n int) ([]byte, error) {
	if b.data == nil {
		return nil, ErrClosed
	} else if b.pos >= len(*b.data) {
		return nil, io.EOF
	} else if b.pos+n > len(*b.data) {
		return (*b.data)[b.pos:], io.EOF
	}
	return (*b.data)[b.pos : b.pos+n], nil
}

// Read is an implementation of the io.Reader interface
func (b *ReadWriteMem) Read(p []byte) (int, error) {
	if b.data == nil {
		return 0, ErrClosed
	} else if b.pos >= len(*b.data) {
		return 0, io.EOF
	}
	n := copy(p, (*b.data)[b.pos:])
	b.pos += n
	return n, nil
}

// ReadByte is an implementation of the io.ByteReader interface
func (b *ReadWriteMem) ReadByte() (byte, error) {
	if b.data == nil {
		return 0, ErrClosed
	} else if b.pos >= len(*b.data) {
		return 0, io.EOF
	}
	c := (*b.data)[b.pos]
	b.pos++
	return c, nil
}

// UnreadByte implements the io.ByteScanner interface
func (b *ReadWriteMem) UnreadByte() error {
	if b.data == nil {
		return ErrClosed
	}
	if b.pos > 0 {
		b.pos--
		return nil
	}
	return ErrInvalidUnreadByte
}

// ReadAt is an implementation of the io.ReaderAt interface
func (b *ReadWriteMem) ReadAt(p []byte, off int64) (int, error) {
	if b.data == nil {
		return 0, ErrClosed
	} else if off >= int64(len(*b.data)) {
		return 0, io.EOF
	}
	return copy(p, (*b.data)[off:]), nil
}

// WriteTo is an implementation of the io.WriterTo interface
func (b *ReadWriteMem) WriteTo(f io.Writer) (int64, error) {
	if b.data == nil {
		return 0, ErrClosed
	} else if b.pos >= len(*b.data) {
		return 0, io.EOF
	}
	n, err := f.Write((*b.data)[b.pos:])
	b.pos = len(*b.data)
	return int64(n), err
}

// Errors
var (
	ErrInvalidUnreadByte = errors.New("invalid UnreadByte, no bytes read")
)
