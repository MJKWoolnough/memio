// Package memio implements Read, Write, Seek, Close and other io methods for a byte slice.
package memio

import "io"

const (
	seekSet = iota
	seekCurr
	seekEnd
)

// Closed is an error returned when trying to perform an operation after using Close().
type Closed struct{}

func (Closed) Error() string {
	return "operation not permitted when closed"
}

type readMem struct {
	data []byte
	pos  int
}

// Open uses a byte slice for reading. Implements io.Reader, io.Seeker,
// io.Closer, io.ReaderAt, io.ByteReader and io.WriterTo.
func Open(data []byte) io.Reader {
	return &readMem{data, 0}
}

func (b *readMem) Read(p []byte) (int, error) {
	if b.data == nil {
		return 0, &Closed{}
	} else if b.pos >= len(b.data) {
		return 0, io.EOF
	}
	n := copy(p, b.data[b.pos:])
	b.pos += n
	return n, nil
}

func (b *readMem) ReadByte() (byte, error) {
	if b.data == nil {
		return 0, &Closed{}
	} else if b.pos >= len(b.data) {
		return 0, io.EOF
	}
	c := b.data[b.pos]
	b.pos++
	return c, nil
}

func (b *readMem) Seek(offset int64, whence int) (int64, error) {
	if b.data == nil {
		return 0, &Closed{}
	}
	switch whence {
	case seekSet:
		b.pos = int(offset)
	case seekCurr:
		b.pos += int(offset)
	case seekEnd:
		b.pos = len(b.data) - int(offset)
	}
	if b.pos < 0 {
		b.pos = 0
	}
	return int64(b.pos), nil
}

func (b *readMem) Close() error {
	b.data = nil
	return nil
}

func (b *readMem) ReadAt(p []byte, off int64) (int, error) {
	if b.data == nil {
		return 0, &Closed{}
	} else if off >= int64(len(b.data)) {
		return 0, io.EOF
	}
	return copy(p, b.data[off:]), nil
}

func (b *readMem) WriteTo(f io.Writer) (int64, error) {
	if b.data == nil {
		return 0, &Closed{}
	} else if b.pos >= len(b.data) {
		return 0, io.EOF
	}
	n, err := f.Write(b.data[b.pos:])
	b.pos = len(b.data)
	return int64(n), err
}

type writeMem struct {
	data *[]byte
	pos  int
}

// Create uses a byte slice for writing. Implements io.Writer, io.Seeker,
// io.Closer, io.WriterAt, io.ByteWriter and io.ReaderFrom.
func Create(data *[]byte) io.Writer {
	return &writeMem{data, 0}
}

func (b *writeMem) Write(p []byte) (int, error) {
	if b.data == nil {
		return 0, &Closed{}
	}
	b.setSize(b.pos + len(p))
	n := copy((*b.data)[b.pos:], p)
	b.pos += n
	return n, nil
}

func (b *writeMem) WriteAt(p []byte, off int64) (int, error) {
	if b.data == nil {
		return 0, &Closed{}
	}
	b.setSize(int(off) + len(p))
	return copy((*b.data)[off:], p), nil
}

func (b *writeMem) WriteByte(c byte) error {
	if b.data == nil {
		return &Closed{}
	}
	b.setSize(b.pos + 1)
	(*b.data)[b.pos] = c
	b.pos++
	return nil
}

func (b *writeMem) ReadFrom(f io.Reader) (int64, error) {
	if b.data == nil {
		return 0, &Closed{}
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

func (b *writeMem) Seek(offset int64, whence int) (int64, error) {
	if b.data == nil {
		return 0, &Closed{}
	}
	switch whence {
	case seekSet:
		b.pos = int(offset)
	case seekCur:
		b.pos += int(offset)
	case seekEnd:
		b.pos = len(*b.data) - int(offset)
	}
	if b.pos < 0 {
		b.pos = 0
	}
	return int64(b.pos), nil
}

func (b *writeMem) Close() error {
	b.data = nil
	return nil
}

func (b *writeMem) setSize(end int) {
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
