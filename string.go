package memio

import "io"

// String grants a string Read-Only methods.
type String string

// Read satisfies the io.Reader interface
func (s *String) Read(p []byte) (int, error) {
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
func (s *String) WriteTo(w io.Writer) (int64, error) {
	n, err := io.WriteString(w, string(*s))
	*s = (*s)[n:]
	return int64(n), err
}

// ReadByte satisfies the io.ByteReader interface
func (s *String) ReadByte() (byte, error) {
	if len(*s) == 0 {
		return 0, io.EOF
	}
	b := (*s)[0]
	*s = (*s)[1:]
	return b, nil
}

// Peek reads the next n bytes without advancing the position
func (s *String) Peek(n int) ([]byte, error) {
	if n > len(*s) {
		return []byte(*s), io.EOF
	}
	return []byte((*s)[:n]), nil
}

// Close satisfies the io.Closer interface
func (s *String) Close() error {
	*s = ""
	return nil
}
