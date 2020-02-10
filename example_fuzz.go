// +build gofuzz

package fuzzme

import "fmt"

func Fuzz(data []byte) int {
	// Create fuzzed input from the fuzzed data.
	f := New(data)
	i := f.Int()

	// Test the function with the fuzzed input.
	example(i)
	return 0
}

// example is the function that is being tested.
func example(i int) {
	if i%12313 == 0 {
		panic(fmt.Sprintf("found it: %d", i))
	}
}
