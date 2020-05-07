package json

import (
	"bytes"
	"fmt"
	"io"
)

type aReadCloser struct {
	r io.Reader
}

func (a *aReadCloser) Read(b []byte) (n int, err error) {
	return a.r.Read(b)
}

func (a *aReadCloser) Close() error {
	return nil
}

func ExampleUnmarshal() {

	reader := &aReadCloser{r: bytes.NewReader([]byte("[\"Hello\", \"world\"]"))}

	var data []string

	Unmarshal(reader, &data)

	fmt.Println(data)
	// Output: [Hello world]
}
