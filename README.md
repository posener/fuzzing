# fuzzing

[![Build Status](https://travis-ci.org/posener/fuzzing.svg?branch=master)](https://travis-ci.org/posener/fuzzing)
[![codecov](https://codecov.io/gh/posener/fuzzing/branch/master/graph/badge.svg)](https://codecov.io/gh/posener/fuzzing)
[![GoDoc](https://godoc.org/github.com/posener/fuzzing?status.svg)](http://godoc.org/github.com/posener/fuzzing)
[![goreadme](https://goreadme.herokuapp.com/badge/posener/fuzzing.svg)](https://goreadme.herokuapp.com)

Package fuzzing enables easy fuzzing with [go-fuzz](https://github.com/dvyukov/go-fuzz).

The `Fuzz` object provides functions for generating consistent Go primitive values from a given
fuzzed bytes slice. The generated values are promised to be consistent from identical slices.
They are also correlated to the given fuzzed slice to enable fuzzing exploration.

For an example on how to use this library with go-fuzz, see [./example_fuzz.go](./example_fuzz.go)

#### Examples

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

Created by [goreadme](https://github.com/apps/goreadme)