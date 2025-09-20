# memio

[![CI](https://github.com/MJKWoolnough/memio/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/memio/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/memio.svg)](https://pkg.go.dev/vimagination.zapto.org/memio)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/memio)](https://goreportcard.com/report/vimagination.zapto.org/memio)

--
    import "vimagination.zapto.org/memio"

Package memio implements Read, Write, Seek, Close and other io methods for a byte slice.

## Highlights

 - `memio.Buffer`: a slice that implements many IO interfaces. It advances the length of the slice as bytes are read, and moves the start of the slice as bytes are read. Some of the interfaces implemented are:
   - `io.Reader`
   - `io.ReaderFrom`
   - `io.ByteReader`
   - `io.RuneReader`
   - `io.Writer`
   - `io.WriterTo`
   - `io.ByteWriter`
   - `io.RuneWriter`
   - `io.Closer`
   - `io.ReaderAt`
   - `io.WriterAt`
   - & more.
 - `memio.LimitedBuffer`: similar to `memio.Buffer`, but will not grow beyond it's capacity.
 - `memio.ReadMem`: a wrapper around `bytes.Reader` that also implements `io.Closer` and a `Peek` method.
 - `memio.WriteMem`: a more compatible version of `memio.Buffer` that doesn't forget read bytes.

## Usage

```go
package main

import (
	"errors"
	"fmt"
	"io"

	"vimagination.zapto.org/memio"
)

func main() {
	var buf memio.Buffer

	buf.WriteString("Hello, world!")

	data := make([]byte, 5)

	buf.Read(data)

	fmt.Printf("%s", data)
	// Output: Hello
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/memio
