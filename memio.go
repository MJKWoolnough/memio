// Copyright (c) 2013 - Michael Woolnough <michael.woolnough@gmail.com>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// MemIO implements Read, Write, Seek, Close and other io methods for a byte slice.
package memio

import "io"

const (
	SEEK_SET int = iota
	SEEK_CURR
	SEEK_END
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

// Use a byte slice for reading. Implements io.Reader, io.Seeker, io.Closer, io.ReaderAt, io.ByteReader and io.WriterTo.
func Open(data []byte) *readMem {
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
	case SEEK_SET:
		b.pos = int(offset)
	case SEEK_CURR:
		b.pos += int(offset)
	case SEEK_END:
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

// Use a byte slice for writing. Implements io.Writer, io.Seeker, io.Closer, io.WriterAt, io.ByteWriter and io.ReaderFrom.
func Create(data *[]byte) *writeMem {
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
	case SEEK_SET:
		b.pos = int(offset)
	case SEEK_CURR:
		b.pos += int(offset)
	case SEEK_END:
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
