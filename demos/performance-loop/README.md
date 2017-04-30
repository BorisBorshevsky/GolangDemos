### Golang 1.8.1



`go test -bench=. -benchmem`

```
BenchmarkInlineN-4          1000           1363190 ns/op         9247346 B/op         30 allocs/op
BenchmarkOutsideN-4         1000           1387100 ns/op         9247347 B/op         30 allocs/op
BenchmarkInsideN-4          1000           1390054 ns/op         9247344 B/op         30 allocs/op
BenchmarkPointerN-4          300           4922183 ns/op         6254328 B/op     100030 allocs/op


```