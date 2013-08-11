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

// Memio implements Read, Write, Seek and Close methods for a byte slice.
package memio

import "io"

type Closed struct{}

func (_ Closed) Error() string {
	return "operation not permitted when closed"
}

type readMem struct {
	data []byte
	pos  int
}

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

func (b *readMem) Seek(offset int64, whence int) (int64, error) {
	if b.data == nil {
		return 0, &Closed{}
	}
	switch whence {
	case 0:
		b.pos = int(offset)
	case 1:
		b.pos += int(offset)
	case 2:
		b.pos = len(b.data) - int(offset) //-1?
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

type writeMem struct {
	data *[]byte
	pos  int
}

func Create(data *[]byte) *writeMem {
	return &writeMem{data, 0}
}

func (b *writeMem) Write(p []byte) (int, error) {
	if b.data == nil {
		return 0, &Closed{}
	}
	if end := b.pos + len(p); end > len(*b.data) {
		t := make([]byte, end)
		copy(t, *b.data)
		*b.data = t
	}
	n := copy((*b.data)[b.pos:], p)
	b.pos += n
	return n, nil
}

func (b *writeMem) Seek(offset int64, whence int) (int64, error) {
	if b.data == nil {
		return 0, &Closed{}
	}
	switch whence {
	case 0:
		b.pos = int(offset)
	case 1:
		b.pos += int(offset)
	case 2:
		b.pos = len(*b.data) - int(offset) //-1?
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
