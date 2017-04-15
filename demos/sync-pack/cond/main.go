package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	n := 5

	c.L.Lock()

	go func() {
		for i := 0; i < n; i++ {
			fmt.Printf("before wait %v \n", i)
			c.Wait()
			fmt.Printf("after wait %v \n", i)
		}
	}()

	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Printf("before signal %v \n", i)
		c.Signal()
		fmt.Printf("after signal %v \n", i)
	}

	time.Sleep(time.Second)

	fmt.Println("we are done")
}
