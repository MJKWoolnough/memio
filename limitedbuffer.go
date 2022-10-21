package memio

import (
	"io"
	"unicode/utf8"
)

// LimitedBuffer grants a byte slice very straightforward IO methods, limiting
// writing to the capacity of the slice
type LimitedBuffer []byte

// Read satisfies the io.Reader interface
func (s *LimitedBuffer) Read(p []byte) (int, error) {
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

// ReadAt satisfies the io.ReaderAt interface
//
// Care should be taken when used in conjunction with any other Read* calls as
// they will alter the start point of the buffer
func (s *LimitedBuffer) ReadAt(p []byte, off int64) (int, error) {
	n := copy(p, (*s)[off:])
	if n < len(p) {
		return n, io.ErrUnexpectedEOF
	}
	return n, nil
}

// WriteTo satisfies the io.WriterTo interface
func (s *LimitedBuffer) WriteTo(w io.Writer) (int64, error) {
	if len(*s) == 0 {
		return 0, io.EOF
	}
	n, err := w.Write(*s)
	*s = (*s)[n:]
	return int64(n), err
}

// Write satisfies the io.Writer interface
func (s *LimitedBuffer) Write(p []byte) (int, error) {
	var err error
	if left := cap(*s) - len(*s); len(p) > left {
		p = p[:left]
		err = io.ErrShortBuffer
	}
	*s = append(*s, p...)
	return len(p), err
}

// WriteAt satisfies the io.WriterAt interface
func (s *LimitedBuffer) WriteAt(p []byte, off int64) (int, error) {
	if off+int64(len(p)) >= int64(cap(p)) {
		return 0, io.ErrShortWrite
	}
	n := copy((*s)[off:cap(*s)], p)
	if n < len(p) {
		return n, io.ErrShortWrite
	}
	return n, nil
}

// WriteString writes a string to the buffer without casting to a byte slice
func (s *LimitedBuffer) WriteString(str string) (int, error) {
	var err error
	if left := cap(*s) - len(*s); len(str) > left {
		str = str[:left]
		err = io.ErrShortBuffer
	}
	*s = append(*s, str...)
	return len(str), err
}

// ReadFrom satisfies the io.ReaderFrom interface
func (s *LimitedBuffer) ReadFrom(r io.Reader) (int64, error) {
	var n int64
	for len(*s) < cap(*s) {
		m, err := r.Read((*s)[len(*s):cap(*s)])
		*s = (*s)[:len(*s)+m]
		n += int64(m)
		if err != nil {
			if err == io.EOF {
				break
			}
			return n, err
		}
	}
	return n, nil
}

// ReadByte satisfies the io.ByteReader interface
func (s *LimitedBuffer) ReadByte() (byte, error) {
	if len(*s) == 0 {
		return 0, io.EOF
	}
	b := (*s)[0]
	*s = (*s)[1:]
	return b, nil
}

// ReadRune satisfies the io.RuneReader interface
func (s *LimitedBuffer) ReadRune() (rune, int, error) {
	if len(*s) == 0 {
		return 0, 0, io.EOF
	}
	r, n := utf8.DecodeRune(*s)
	*s = (*s)[n:]
	return r, n, nil
}

// WriteByte satisfies the io.ByteWriter interface
func (s *LimitedBuffer) WriteByte(b byte) error {
	if len(*s) == cap(*s) {
		return io.EOF
	}
	*s = append(*s, b)
	return nil
}

// Peek reads the next n bytes without advancing the position
func (s *LimitedBuffer) Peek(n int) ([]byte, error) {
	if *s == nil {
		return nil, ErrClosed
	} else if n > len(*s) {
		return *s, io.EOF
	}
	return (*s)[:n], nil
}

// Close satisfies the io.Closer interface
func (s *LimitedBuffer) Close() error {
	*s = nil
	return nil
}
