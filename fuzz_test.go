package fuzzing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	f := New([]byte{1, 2, 3})
	i := f.Int()
	fmt.Println(i)
	// Output: 3851489450890114710
}

var seedBytes = []byte{0, 0, 0, 0, 0, 0, 0, 1}

func new(input ...byte) *Fuzz {
	return New(append(seedBytes, input...))
}

func TestFuzzRand(t *testing.T) {
	t.Parallel()

	f := new()

	assert.Equal(t, 5980212987775051087, f.Int())
	assert.Equal(t, 1603104512986455410, f.Int())
}

func TestFuzzRead(t *testing.T) {
	t.Parallel()

	f := new(9, 10)

	// 9th byte.
	assert.Equal(t, []byte{9}, f.Bytes(1))
	// 10th byte.
	assert.Equal(t, []byte{10}, f.Bytes(1))
	// From seed.
	assert.Equal(t, []byte{82}, f.Bytes(1))
}

func TestFuzzRead_combine(t *testing.T) {
	t.Parallel()

	f := new(9)

	assert.Equal(t, []byte{9 /* 9th byte */, 82 /* From seed */}, f.Bytes(2))
}

func TestFuncs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		data       []byte
		wantInt    int
		wantUint   uint
		wantInt64  int64
		wantInt63  int64
		wantUint64 uint64
		wantInt32  int32
		wantInt31  int32
		wantUint32 uint32
		wantBool   bool
		wantRune   rune
	}{
		{
			data:       []byte{0, 0, 0, 0, 0, 0, 0, 1},
			wantInt:    1,
			wantUint:   1,
			wantInt64:  1,
			wantInt63:  0,
			wantUint64: 1,
			wantInt32:  0,
			wantInt31:  0,
			wantUint32: 0,
			wantBool:   false,
			wantRune:   '\x00',
		},
		{
			data:       []byte{0, 0, 0, 0, 0, 0, 2, 3},
			wantInt:    515,
			wantUint:   515,
			wantInt64:  515,
			wantInt63:  257,
			wantUint64: 515,
			wantInt32:  0,
			wantInt31:  0,
			wantUint32: 0,
			wantBool:   false,
			wantRune:   '\x00',
		},
		{
			data:       []byte{255, 255, 255, 255, 255, 255, 255, 255},
			wantInt:    -1,
			wantUint:   0xffffffffffffffff,
			wantInt64:  -1,
			wantInt63:  0x7fffffffffffffff,
			wantUint64: 0xffffffffffffffff,
			wantInt32:  -1,
			wantInt31:  0x7fffffff,
			wantUint32: 0xffffffff,
			wantBool:   true,
			wantRune:   '�',
		},
		{
			data:       nil,
			wantInt:    5980212987775051087,
			wantUint:   5980212987775051087,
			wantInt64:  5980212987775051087,
			wantInt63:  2990106493887525543,
			wantUint64: 5980212987775051087,
			wantInt32:  1392376839,
			wantInt31:  696188419,
			wantUint32: 1392376839,
			wantBool:   false,
			wantRune:   'R',
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.data), func(t *testing.T) {
			assert.Equal(t, tt.wantInt, new(tt.data...).Int(), "int")
			assert.Equal(t, tt.wantUint, new(tt.data...).Uint(), "uint")
			assert.Equal(t, tt.wantInt64, new(tt.data...).Int64(), "int64")
			assert.Equal(t, tt.wantInt63, new(tt.data...).Int63(), "int63")
			assert.Equal(t, tt.wantUint64, new(tt.data...).Uint64(), "uint64")
			assert.Equal(t, tt.wantInt32, new(tt.data...).Int32(), "int32")
			assert.Equal(t, tt.wantInt31, new(tt.data...).Int31(), "int31")
			assert.Equal(t, tt.wantUint32, new(tt.data...).Uint32(), "uint32")
			assert.Equal(t, tt.wantBool, new(tt.data...).Bool(), "bool")
			assert.Equal(t, tt.wantRune, new(tt.data...).Rune(), "rune")
		})
	}
}

func TestFuzzString(t *testing.T) {
	t.Parallel()

	f := new('a', 'b', 'c')

	// 9th byte as a rune.
	assert.Equal(t, "a", f.String(1))
	// 10th, 11th bytes as runes.
	assert.Equal(t, "bc", f.String(2))
	// From seed.
	assert.Equal(t, "R", f.String(1))
}

func TestNewPanics(t *testing.T) {
	assert.Panics(t, func() { New(nil) })
}
