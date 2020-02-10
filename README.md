# fuzzing

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
