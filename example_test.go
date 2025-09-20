package memio_test

import (
	"errors"
	"fmt"
	"io"

	"vimagination.zapto.org/memio"
)

// ExampleBuffer demonstrates writing and reading using Buffer.
func ExampleBuffer() {
	var buf memio.Buffer

	buf.WriteString("Hello, world!")

	data := make([]byte, 5)

	buf.Read(data)

	fmt.Printf("%s", data)
	// Output: Hello
}

// ExampleLimitedBuffer shows how LimitedBuffer prevents writes beyond capacity.
func ExampleLimitedBuffer() {
	// Preallocate 5 bytes
	data := make([]byte, 5)
	lb := memio.LimitedBuffer(data)

	// Write exactly 5 bytes
	lb.WriteString("12345")

	// Attempting to write more will not extend capacity
	n, err := lb.WriteString("678")
	fmt.Println(n, errors.Is(err, io.ErrShortBuffer))
	// Output: 0 true
}
