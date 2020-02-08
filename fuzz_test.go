package fuzzing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	f := New([]byte{1, 2, 3})
	i := f.SignedInt()
	fmt.Println(i)
	// Output: -2781883647095912858
}

var seedBytes = []byte{1, 2, 3, 4, 5, 6, 7, 8}

func new(input ...byte) *Fuzz {
	return New(append(seedBytes, input...))
}

func TestFuzzRand(t *testing.T) {
	t.Parallel()

	f := new(9)

	assert.Equal(t, 324259173170675712, f.Int())
	assert.Equal(t, 6024040606121685663, f.Int())
}

func TestFuzzRead(t *testing.T) {
	t.Parallel()

	f := new(9, 10)

	assert.Equal(t, []byte{9}, f.Bytes(1))
	assert.Equal(t, []byte{10}, f.Bytes(1))
	assert.Equal(t, []byte{159}, f.Bytes(1)) // From rand source.
}

func TestFuzzRead_combine(t *testing.T) {
	t.Parallel()

	f := new(9)

	assert.Equal(t, []byte{9, 159}, f.Bytes(2))
}

func TestFuzzBool(t *testing.T) {
	t.Parallel()

	assert.False(t, new(0).Bool())
	assert.True(t, new(1).Bool())

	assert.True(t, new().Bool()) // From rand source.
}

func TestSignedInt(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 36311929895191428, new(1, 2, 3, 4, 5, 6, 7, 8, 0).SignedInt())
	assert.Equal(t, -36311929895191428, new(1, 2, 3, 4, 5, 6, 7, 8, 1).SignedInt())

	assert.Equal(t, 6024040606121685663, new().SignedInt()) // From rand source.
}

func TestSignedInt64(t *testing.T) {
	t.Parallel()
	assert.Equal(t, int64(36311929895191428), new(1, 2, 3, 4, 5, 6, 7, 8, 0).SignedInt64())
	assert.Equal(t, int64(-36311929895191428), new(1, 2, 3, 4, 5, 6, 7, 8, 1).SignedInt64())

	assert.Equal(t, int64(6024040606121685663), new().SignedInt64()) // From rand source.
}
