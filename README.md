# memio
--
    import "github.com/MJKWoolnough/memio"

MemIO implements Read, Write, Seek and Close methods for a byte slice.

## Usage

#### func  Create

```go
func Create(data *[]byte) *writeMem
```
Use a byte slice for writing. Implements io.Writer, io.Seeker and io.Closer.

#### func  Open

```go
func Open(data []byte) *readMem
```
Use a byte slice for reading. Implements io.Reader, io.Seeker and io.Closer.

#### type Closed

```go
type Closed struct{}
```

Closed is an error returned when trying to perform an operation after using
Close().

#### func (Closed) Error

```go
func (_ Closed) Error() string
```
