# memio
--
    import "github.com/MJKWoolnough/memio"

Package memio implements Read, Write, Seek, Close and other io methods for a byte slice.

## Usage

```go
var ErrClosed = errors.New("operation not permitted when closed")
```
ErrClosed is an error returned when trying to perform an operation after using
Close().

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

#### func (*ReadWriteMem) WriteTo

```go
func (b *ReadWriteMem) WriteTo(f io.Writer) (int64, error)
```
WriteTo is an implementation of the io.WriterTo interface

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
Close is and implementation of the io.Closer interface

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
