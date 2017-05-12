package performance_loop

import "testing"

var n = 100000

func benchmarkInline(c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		loopInline(c)
	}
}

func benchmarkInside(c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		loopInside(c)
	}
}

func benchmarkOutside(c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		loopOutside(c)
	}
}

func benchmarkPointer(c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		loopPoiner(c)
	}
}

func BenchmarkInlineN(b *testing.B) {
	benchmarkInline(n, b)
}

func BenchmarkOutsideN(b *testing.B) {
	benchmarkOutside(n, b)
}

func BenchmarkInsideN(b *testing.B) {
	benchmarkInside(n, b)
}

func BenchmarkPointerN(b *testing.B) {
	benchmarkPointer(n, b)
}
