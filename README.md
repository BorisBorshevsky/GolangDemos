# GolangDemos


```
go test -bench=. -benchmem
BenchmarkMutex200-4                  300           5006809 ns/op              83 B/op          3 allocs/op
BenchmarkMutex500-4                   50          31973402 ns/op            1165 B/op         10 allocs/op
BenchmarkMutex1000-4                  10         137232301 ns/op            6976 B/op         79 allocs/op
BenchmarkMutex3000-4                   1        1218393377 ns/op         1204064 B/op       4600 allocs/op
BenchmarkAtomic200-4                2000            795079 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic500-4                 300           5322851 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic1000-4                100          21457508 ns/op              24 B/op          2 allocs/op
BenchmarkAtomic3000-4                 10         191115567 ns/op            1315 B/op          5 allocs/op
```

