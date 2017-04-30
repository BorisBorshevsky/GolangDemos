### Golang 1.8.1

`go test -bench=. -benchmem`

```
BenchmarkMutex200-4                  300           5106701 ns/op              57 B/op          3 allocs/op
BenchmarkMutex500-4                   50          32205556 ns/op            1058 B/op         10 allocs/op
BenchmarkMutex1000-4                  10         132542591 ns/op            7739 B/op         83 allocs/op
BenchmarkMutex3000-4                   1        1286181702 ns/op         1185504 B/op       4494 allocs/op
BenchmarkAtomic200-4                2000            534363 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic500-4                 300           5479963 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic1000-4                100          22155160 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic3000-4                 10         200612616 ns/op              25 B/op          2 allocs/op

```