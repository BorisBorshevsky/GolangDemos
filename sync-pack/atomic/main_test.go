package main

import "testing"

func benchmarkAtomic(c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		atomicRun(c)
	}
}

func benchmarkMutex(c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		mutexRun(c)
	}
}

func benchmarkSemaphore(c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		semaphoreRun(c)
	}
}

func BenchmarkMutex200(b *testing.B) {
	benchmarkMutex(200, b)

}

func BenchmarkMutex500(b *testing.B) {
	benchmarkMutex(500, b)
}

func BenchmarkMutex1000(b *testing.B) {
	benchmarkMutex(1000, b)
}

func BenchmarkMutex3000(b *testing.B) {
	benchmarkMutex(3000, b)
}
func BenchmarkAtomic200(b *testing.B) {
	benchmarkAtomic(200, b)

}

func BenchmarkAtomic500(b *testing.B) {
	benchmarkAtomic(500, b)
}

func BenchmarkAtomic1000(b *testing.B) {
	benchmarkAtomic(1000, b)
}

func BenchmarkAtomic3000(b *testing.B) {
	benchmarkAtomic(3000, b)
}

func BenchmarkSem200(b *testing.B) {
	benchmarkSemaphore(200, b)

}

func BenchmarkSem500(b *testing.B) {
	benchmarkSemaphore(500, b)
}

func BenchmarkSem1000(b *testing.B) {
	benchmarkSemaphore(1000, b)
}

func BenchmarkSem3000(b *testing.B) {
	benchmarkSemaphore(3000, b)
}

func TestTrack(t *testing.T) {
	Track(100)
}
