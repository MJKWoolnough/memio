# memio
--
    import "github.com/MJKWoolnough/memio"

Package memio implements Read, Write, Seek, Close and other io methods for a
byte slice.

## Usage

```go
var ErrClosed = errors.New("operation not permitted when closed")
```
ErrClosed is an error returned when trying to perform an operation after using
Close().

```go
var (
	ErrInvalidUnreadByte = errors.New("invalid UnreadByte, no bytes read")
)
```

#### type Buffer

```go
type Buffer []byte
```

Buffer grants a byte slice very straightforward IO methods.

#### func (*Buffer) Close

```go
func (s *Buffer) Close() error
```
Close satisfies the io.Closer interface

#### func (*Buffer) Peek

```go
func (s *Buffer) Peek(n int) ([]byte, error)
```
Peek reads the next n bytes without advancing the position

#### func (*Buffer) Read

```go
func (s *Buffer) Read(p []byte) (int, error)
```
Read satisfies the io.Reader interface

#### func (*Buffer) ReadByte

```go
func (s *Buffer) ReadByte() (byte, error)
```
ReadByte satisfies the io.ByteReader interface

#### func (*Buffer) ReadRune

```go
func (s *Buffer) ReadRune() (rune, int, error)
```
ReadRune satisfies the io.RuneReader interface

#### func (*Buffer) Write

```go
func (s *Buffer) Write(p []byte) (int, error)
```
Write satisfies the io.Writer interface

#### func (*Buffer) WriteByte

```go
func (s *Buffer) WriteByte(b byte) error
```
WriteByte satisfies the io.ByteWriter interface

#### func (*Buffer) WriteString

```go
func (s *Buffer) WriteString(str string) (int, error)
```
WriteString writes a string to the buffer without casting to a byte slice

#### func (*Buffer) WriteTo

```go
func (s *Buffer) WriteTo(w io.Writer) (int64, error)
```
WriteTo satisfies the io.WriterTo interface

#### type ReadMem

```go
type ReadMem struct {
	*bytes.Reader
}
```

ReadMem holds a byte slice that can be used for many io interfaces

#### func  Open

```go
func Open(data []byte) ReadMem
```
Open uses a byte slice for reading. Implements io.Reader, io.Seeker, io.Closer,
io.ReaderAt, io.ByteReader and io.WriterTo.

#### func (ReadMem) Close

```go
func (ReadMem) Close() error
```
Close is a no-op func the lets ReadMem implement interfaces that require a Close
method

#### func (ReadMem) Peek

```go
func (r ReadMem) Peek(n int) ([]byte, error)
```
Peek reads the next n bytes without advancing the position

#### type ReadWriteMem

```go
type ReadWriteMem struct {
	WriteMem
}
```

ReadWriteMem is a combination of both the ReadMem and WriteMem types, allowing
both all reads and writes to the same underlying byte slice.

#### func  OpenMem

```go
func OpenMem(data *[]byte) *ReadWriteMem
```
OpenMem uses a byte slice for reading and writing. Implements io.Reader,
io.Writer, io.Seeker, io.ReaderAt, io.ByteReader, io.WriterTo, io.WriterAt,
io.ByteWriter and io.ReaderFrom.

#### func (*ReadWriteMem) Peek

```go
func (b *ReadWriteMem) Peek(n int) ([]byte, error)
```
Peek reads the next n bytes without advancing the position

#### func (*ReadWriteMem) Read

```go
func (b *ReadWriteMem) Read(p []byte) (int, error)
```
Read is an implementation of the io.Reader interface

#### func (*ReadWriteMem) ReadAt

```go
func (b *ReadWriteMem) ReadAt(p []byte, off int64) (int, error)
```
ReadAt is an implementation of the io.ReaderAt interface

#### func (*ReadWriteMem) ReadByte

```go
func (b *ReadWriteMem) ReadByte() (byte, error)
```
ReadByte is an implementation of the io.ByteReader interface

#### func (*ReadWriteMem) UnreadByte

```go
func (b *ReadWriteMem) UnreadByte() error
```
UnreadByte implements the io.ByteScanner interface

#### func (*ReadWriteMem) WriteTo

```go
func (b *ReadWriteMem) WriteTo(f io.Writer) (int64, error)
```
WriteTo is an implementation of the io.WriterTo interface

#### type String

```go
type String string
```

String grants a string Read-Only methods.

#### func (*String) Close

```go
func (s *String) Close() error
```
Close satisfies the io.Closer interface

#### func (*String) Peek

```go
func (s *String) Peek(n int) ([]byte, error)
```
Peek reads the next n bytes without advancing the position

#### func (*String) Read

```go
func (s *String) Read(p []byte) (int, error)
```
Read satisfies the io.Reader interface

#### func (*String) ReadByte

```go
func (s *String) ReadByte() (byte, error)
```
ReadByte satisfies the io.ByteReader interface

#### func (*String) ReadRune

```go
func (s *String) ReadRune() (rune, int, error)
```
ReadRune satisfies the io.RuneReader interface

#### func (*String) WriteTo

```go
func (s *String) WriteTo(w io.Writer) (int64, error)
```
WriteTo satisfies the io.WriterTo interface

#### type WriteMem

```go
type WriteMem struct {
}
```

WriteMem holds a pointer to a byte slice and allows numerous io interfaces to be
used with it.

#### func  Create

```go
func Create(data *[]byte) *WriteMem
```
Create uses a byte slice for writing. Implements io.Writer, io.Seeker,
io.Closer, io.WriterAt, io.ByteWriter and io.ReaderFrom.

#### func (*WriteMem) Close

```go
func (b *WriteMem) Close() error
```
Close is an implementation of the io.Closer interface

#### func (*WriteMem) ReadFrom

```go
func (b *WriteMem) ReadFrom(f io.Reader) (int64, error)
```
ReadFrom is an implementation of the io.ReaderFrom interface

#### func (*WriteMem) Seek

```go
func (b *WriteMem) Seek(offset int64, whence int) (int64, error)
```
Seek is an implementation of the io.Seeker interface

#### func (*WriteMem) Truncate

```go
func (b *WriteMem) Truncate(s int64) error
```
Truncate changes the length of the byte slice to the given amount

#### func (*WriteMem) Write

```go
func (b *WriteMem) Write(p []byte) (int, error)
```
Write is an implementation of the io.Writer interface

#### func (*WriteMem) WriteAt

```go
func (b *WriteMem) WriteAt(p []byte, off int64) (int, error)
```
WriteAt is an implementation of the io.WriterAt interface

#### func (*WriteMem) WriteByte

```go
func (b *WriteMem) WriteByte(c byte) error
```
WriteByte is an implementation of the io.WriteByte interface

#### func (*WriteMem) WriteString

```go
func (b *WriteMem) WriteString(s string) (int, error)
```
WriteString writes a string to the underlying memory
