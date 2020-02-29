// Package fuzzing enables easy fuzzing with (go-fuzz) https://github.com/dvyukov/go-fuzz.
//
// The `Fuzz` object provides functions for generating consistent Go primitive values from a given
// fuzzed bytes slice. The generated values are promised to be consistent from identical slices.
// They are also correlated to the given fuzzed slice to enable fuzzing exploration.
//
// For an example on how to use this library with go-fuzz, see ./example_fuzz.go
// In order to test the example, run in the project directory:
//
// 	$ go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
// 	$ go-fuzz-build
// 	$ go-fuzz -testoutput
package fuzzing

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"math/rand"
)

// Fuzz is a fuzzing helper object. It provides functions for generating consistent Go primitive
// values from a given fuzzed bytes slice. The generated values are promised to be consistent from
// a identical slices. They are also correlated to the given fuzzed slice to enable fuzzing
// exploration.
type Fuzz struct {
	// reader reads bytes from the 9th byte of the input.
	reader *bytes.Reader
	// fallback is a buffered reader from a random source that is created from the first 8 bytes of
	// the input as a seed.
	fallback *bufio.Reader
}

// New returns a Fuzz.
func New(data []byte) *Fuzz {
	if len(data) <= 0 {
		panic("data must be non empty bytes")
	}
	reader := bytes.NewReader(data)

	var source [8]byte
	_, err := reader.Read(source[:])
	if err != nil && err != io.EOF {
		panic(err)
	}
	seed := int64(binary.BigEndian.Uint64(source[:]))

	return &Fuzz{
		reader:   reader,
		fallback: bufio.NewReader(rand.New(rand.NewSource(seed))),
	}
}

// Read reads from source. If the source was exhausted, it reads from the random fallback. The
// returned error should always be nil.
func (f *Fuzz) Read(b []byte) (int, error) {
	// Try reading from the byte source.
	n, err := f.reader.Read(b)
	if err == io.EOF {
		// If input bytes was exhausted, return random bytes.
		return f.fallback.Read(b)
	}
	if n < len(b) {
		m, err := f.fallback.Read(b[n:])
		return n + m, err
	}
	return n, err
}

// Bytes consumes n bytes and returns them.
func (f *Fuzz) Bytes(n int) []byte {
	b := make([]byte, n, n)
	_, err := f.Read(b)
	if err != nil {
		panic(err) // Should not happen.
	}
	return b
}

// Uint64 consumes 8 bytes and return their conversion to uint64.
func (f *Fuzz) Uint64() uint64 {
	return binary.BigEndian.Uint64(f.Bytes(8))
}

// Int64 consumes 8 bytes and return their conversion to int64.
func (f *Fuzz) Int64() int64 {
	return int64(f.Uint64())
}

// Int63 consumes 8 bytes and return their conversion to int64 in the range [0, 1<<63).
func (f *Fuzz) Int63() int64 {
	return int64(f.Uint64() >> 1)
}

// Uint32 consumes 4 bytes and return their conversion to uint32.
func (f *Fuzz) Uint32() uint32 {
	return binary.BigEndian.Uint32(f.Bytes(4))
}

// Int32 consumes 4 bytes and return their conversion to int32.
func (f *Fuzz) Int32() int32 {
	return int32(f.Uint32())
}

// Int31 consumes 4 bytes and return their conversion to int32 in the range [0, 1<<31).
func (f *Fuzz) Int31() int32 {
	return int32(f.Uint32() >> 1)
}

// Uint consumes 8 bytes and return their conversion to uint.
func (f *Fuzz) Uint() uint {
	return uint(f.Uint64())
}

// Int consumes 8 bytes and return their conversion to int.
func (f *Fuzz) Int() int {
	return int(f.Int64())
}

// Bool consumes one byte and converts it to a boolean value.
func (f *Fuzz) Bool() bool {
	return f.Bytes(1)[0]&1 == 1
}

// ReadRune consumes a single rune. The returned error should always be nil.
func (f *Fuzz) ReadRune() (r rune, size int, err error) {
	// Try reading rune from the byte source.
	r, size, err = f.reader.ReadRune()
	if err == io.EOF {
		// If input bytes was exhausted, return from the random reader.
		return f.fallback.ReadRune()
	}
	return
}

// Rune is a convenience method around ReadRune.
func (f *Fuzz) Rune() rune {
	r, _, err := f.ReadRune()
	if err != nil {
		panic(err) // Should not happen.
	}
	return r
}

// String consumes n runes and return them as a string.
func (f *Fuzz) String(n int) string {
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		runes[i] = f.Rune()
	}
	return string(runes)
}
