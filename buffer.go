package memio

import "io"

// Buffer grants a byte slice very straightforward IO methods.
// Write methods do not expand the length/capacity of the slice
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
	n, err := w.Write([]byte(*s))
	*s = (*s)[n:]
	return int64(n), err
}

// Write satisfies the io.Writer interface
func (s *Buffer) Write(p []byte) (int, error) {
	*s = append(*s, p...)
	return len(p), nil
}
