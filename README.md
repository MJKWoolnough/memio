# memio
--
    import "github.com/MJKWoolnough/memio"

MemIO implements Read, Write, Seek, Close and other io methods for a byte slice.

## Usage

#### func  Create

```go
func Create(data *[]byte) *writeMem
```
Use a byte slice for writing. Implements io.Writer, io.Seeker, io.Closer,
io.WriterAt, io.ByteWriter and io.ReaderFrom.

#### func  Open

```go
func Open(data []byte) *readMem
```
Use a byte slice for reading. Implements io.Reader, io.Seeker, io.Closer,
io.ReaderAt, io.ByteReader and io.WriterTo.

#### type Closed

```go
type Closed struct{}
```

Closed is an error returned when trying to perform an operation after using
Close().

#### func (Closed) Error

```go
func (Closed) Error() string
```
