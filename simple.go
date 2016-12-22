package memio

import "io"

// Simple grants a byte slice very straightfoward read/write methods.
// Write methods do not expand the length/capacity of the slice
type Simple []byte

// Read satisfies the io.Reader interface
func (s *Simple) Read(p []byte) (int, error) {
	n := copy(p, *s)
	*s = (*s)[n:]
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

// WriteTo satisfies the io.WriterTo interface
func (s *Simple) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(*s))
	*s = (*s)[n:]
	return int64(n), err
}

// Write satisfies the io.Writer interface
func (s *Simple) Write(p []byte) (int, error) {
	n := copy(*s, p)
	*s = (*s)[n:]
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

// ReadFrom satifies the io.ReaderFrom interface
func (s *Simple) ReadFrom(r io.Reader) (int64, error) {
	n, err := io.ReadFull(r, *s)
	*s = (*s)[n:]
	if err == io.EOF {
		err = nil
	}
	return int64(n), err
}
