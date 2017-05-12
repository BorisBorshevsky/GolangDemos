package common

import (
	"log"
	"runtime"
	"time"
)

func TimeTrack(fn func(int) int64, n int, name string) {
	runtime.GC()
	start := time.Now()
	res := fn(n)
	elapsed := time.Since(start)
	log.Printf("function %s took %s, result was %v", name, elapsed, res)
}
