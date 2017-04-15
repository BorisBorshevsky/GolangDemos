package main

import (
	"sync/atomic"

	"sync"

	"github.com/gettaxi/smallDemos/demos/common"
)

func main() {
	Track(1000)
}

func Track(n int) {
	common.TimeTrack(simpleRun, n, "simple")

	common.TimeTrack(mutexRun, n, "mutex")

	common.TimeTrack(atomicRun, n, "atomic")
}

func simpleRun(n int) int64 {
	var a int64 = 0

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				a += 1
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return a
}

func mutexRun(n int) int64 {
	var a int64 = 0

	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				m.Lock()
				a += 1
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return a

}

func atomicRun(n int) int64 {
	var a int64 = 0

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				atomic.AddInt64(&a, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return a
}
