### Golang 1.8.1

`go test -bench=. -benchmem`

```
BenchmarkMutex200-4                  300           4553057 ns/op              87 B/op          3 allocs/op
BenchmarkMutex500-4                   50          32217276 ns/op            1009 B/op         11 allocs/op
BenchmarkMutex1000-4                  10         129875189 ns/op            7420 B/op         83 allocs/op
BenchmarkMutex3000-4                   1        1245086432 ns/op         1209536 B/op       4555 allocs/op
BenchmarkAtomic200-4                2000            808973 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic500-4                 300           5221206 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic1000-4                100          18421669 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic3000-4                 10         179790372 ns/op            1232 B/op          4 allocs/op
BenchmarkSem200-4                    100          12236419 ns/op             328 B/op          5 allocs/op
BenchmarkSem500-4                     20          72778576 ns/op            1969 B/op         25 allocs/op
BenchmarkSem1000-4                     3         335808412 ns/op           23658 B/op        297 allocs/op
BenchmarkSem3000-4                     1        5200870069 ns/op          227024 B/op       2839 allocs/op

```