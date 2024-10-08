package memio

import (
	"io"
	"unicode/utf8"
)

// Buffer grants a byte slice very straightforward IO methods.
type Buffer []byte

// Read satisfies the io.Reader interface.
func (s *Buffer) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	if len(*s) == 0 {
		return 0, io.EOF
	}

	n := copy(p, *s)
	*s = (*s)[n:]

	return n, nil
}

// ReadAt satisfies the io.ReaderAt interface.
//
// Care should be taken when used in conjunction with any other Read* calls as
// they will alter the start point of the buffer.
func (s *Buffer) ReadAt(p []byte, off int64) (int, error) {
	n := copy(p, (*s)[off:])
	if n < len(p) {
		return n, io.EOF
	}

	return n, nil
}

// WriteTo satisfies the io.WriterTo interface.
func (s *Buffer) WriteTo(w io.Writer) (int64, error) {
	if len(*s) == 0 {
		return 0, io.EOF
	}

	n, err := w.Write(*s)
	*s = (*s)[n:]

	return int64(n), err
}

// Write satisfies the io.Writer interface.
func (s *Buffer) Write(p []byte) (int, error) {
	*s = append(*s, p...)
	return len(p), nil
}

// WriteAt satisfies the io.WriteAt interface.
func (s *Buffer) WriteAt(p []byte, off int64) (int, error) {
	l := int64(len(p)) + off
	if int64(cap(*s)) < l {
		t := make([]byte, len(*s), l)

		copy(t, (*s)[:cap(*s)])

		*s = t
	}

	return copy((*s)[off:cap(*s)], p), nil
}

// WriteString writes a string to the buffer without casting to a byte slice.
func (s *Buffer) WriteString(str string) (int, error) {
	*s = append(*s, str...)

	return len(str), nil
}

// ReadFrom satisfies the io.ReaderFrom interface.
func (s *Buffer) ReadFrom(r io.Reader) (int64, error) {
	var n int64

	for {
		if len(*s) == cap(*s) {
			*s = append(*s, 0)[:len(*s)]
		}

		m, err := r.Read((*s)[len(*s):cap(*s)])
		*s = (*s)[:len(*s)+m]
		n += int64(m)

		if err != nil {
			if err == io.EOF {
				return n, nil
			}

			return n, err
		}
	}
}

// ReadByte satisfies the io.ByteReader interface.
func (s *Buffer) ReadByte() (byte, error) {
	if len(*s) == 0 {
		return 0, io.EOF
	}

	b := (*s)[0]
	*s = (*s)[1:]

	return b, nil
}

// ReadRune satisfies the io.RuneReader interface.
func (s *Buffer) ReadRune() (rune, int, error) {
	if len(*s) == 0 {
		return 0, 0, io.EOF
	}

	r, n := utf8.DecodeRune(*s)
	*s = (*s)[n:]

	return r, n, nil
}

// WriteByte satisfies the io.ByteWriter interface.
func (s *Buffer) WriteByte(b byte) error {
	*s = append(*s, b)

	return nil
}

// Peek reads the next n bytes without advancing the position.
func (s *Buffer) Peek(n int) ([]byte, error) {
	if *s == nil {
		return nil, ErrClosed
	} else if n > len(*s) {
		return *s, io.EOF
	}

	return (*s)[:n], nil
}

// Close satisfies the io.Closer interface.
func (s *Buffer) Close() error {
	*s = nil

	return nil
}
