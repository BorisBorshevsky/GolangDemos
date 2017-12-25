package main

import (
	"sync/atomic"

	"github.com/BorisBorshevsky/GolangDemos/common"

	"sync"
)

func main() {
	Track(1000)
}

func Track(n int) {
	//common.TimeTrack(simpleRun, n, "simple")

	common.TimeTrack(mutexRun, n, "mutex")

	common.TimeTrack(atomicRun, n, "atomic")

	common.TimeTrack(semaphoreRun, n, "sem")

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

	m := &sync.Mutex{}
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

func semaphoreRun(n int) int64 {
	var a int64 = 0
	var sem = make(chan int, 1)

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				sem <- 1
				a += 1
				<-sem
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return a
}
