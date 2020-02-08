// Package fuzzing enables easy fuzzing with go-fuzz (https://github.com/dvyukov/go-fuzz).
//
// The `Fuzz` object provides functions for generating consistent Go primitive values from a fuzzed
// a given bytes slice. The generated values are promised to be consistent from identical slices.
// They are also correlated to the given fuzzed slice to enable fuzzing exploration.
//
// For an example on how to use this library with go-fuzz, see ./example_fuzz.go
package fuzzing

import (
	"io"
	"math/rand"

	"github.com/posener/fuzzing/internal/bytesource"
)

// Fuzz is a fuzzing helper object. It provides functions for generating consistent Go primitive
// values from a fuzzed a given bytes slice. The generated values are promised to be consistent from
// a identical slices. They are also correlated to the given fuzzed slice to enable fuzzing
// exploration.
type Fuzz struct {
	*rand.Rand
	source *bytesource.ByteSource
}

// New returns a Fuzz.
func New(data []byte) *Fuzz {
	b := bytesource.New(data)
	return &Fuzz{
		Rand:   rand.New(b),
		source: b,
	}
}

// Read reads from source. If the source was exhausted, it reads from the random fallback.
func (f *Fuzz) Read(b []byte) (int, error) {
	// Try reading from the byte source.
	n, err := f.source.Read(b)
	if err == io.EOF {
		// If input bytes was exhausted, return random bytes.
		return f.Rand.Read(b)
	}
	if n < len(b) {
		m, err := f.Rand.Read(b[n:])
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

// Bool consumes a byte and converts it to a boolean value.
func (f *Fuzz) Bool() bool {
	return f.Bytes(1)[0]&1 == 1
}

// SignedInt consumes an int and sets its sign by consuming another byte.
func (f *Fuzz) SignedInt() int {
	i := f.Int()
	if f.Bool() {
		i = -i
	}
	return i
}

// SignedInt consumes an int64 and sets its sign by consuming another byte.
func (f *Fuzz) SignedInt64() int64 {
	i := f.Int63()
	if f.Bool() {
		i = -i
	}
	return i
}
