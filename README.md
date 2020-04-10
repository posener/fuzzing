# fuzzing

[![codecov](https://codecov.io/gh/posener/fuzzing/branch/master/graph/badge.svg)](https://codecov.io/gh/posener/fuzzing)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/posener/fuzzing)

Package fuzzing enables easy fuzzing with [go-fuzz](https://github.com/dvyukov/go-fuzz).

The `Fuzz` object provides functions for generating consistent Go primitive values from a given
fuzzed bytes slice. The generated values are promised to be consistent from identical slices.
They are also correlated to the given fuzzed slice to enable fuzzing exploration.

For an example on how to use this library with go-fuzz, see [./example_fuzz.go](./example_fuzz.go)
In order to test the example, run in the project directory:

```go
$ go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
$ go-fuzz-build
$ go-fuzz -testoutput
```

## Examples

```golang
f := New([]byte{1, 2, 3})
i := f.Int()
fmt.Println(i)
```

 Output:

```
3851489450890114710
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
