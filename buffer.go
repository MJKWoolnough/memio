package memio

import "io"

// Buffer grants a byte slice very straightforward IO methods.
type Buffer []byte

// Read satisfies the io.Reader interface
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

// WriteTo satisfies the io.WriterTo interface
func (s *Buffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(*s)
	*s = (*s)[n:]
	return int64(n), err
}

// Write satisfies the io.Writer interface
func (s *Buffer) Write(p []byte) (int, error) {
	*s = append(*s, p...)
	return len(p), nil
}

// WriteString writes a string to the buffer without casting to a byte slice
func (s *Buffer) WriteString(str string) (int, error) {
	*s = append(*s, str...)
	return len(str), nil
}

// ReadByte satisfies the io.ByteReader interface
func (s *Buffer) ReadByte() (byte, error) {
	if len(*s) == 0 {
		return 0, io.EOF
	}
	b := (*s)[0]
	*s = (*s)[1:]
	return b, nil
}

// WriteByte satisfies the io.ByteWriter interface
func (s *Buffer) WriteByte(b byte) error {
	*s = append(*s, b)
	return nil
}

// Peek reads the next n bytes without advancing the position
func (s *Buffer) Peek(n int) ([]byte, error) {
	if *s == nil {
		return nil, ErrClosed
	} else if n > len(*s) {
		return *s, io.EOF
	}
	return (*s)[:n], nil
}

// Close satisfies the io.Closer interface
func (s *Buffer) Close() error {
	*s = nil
	return nil
}
