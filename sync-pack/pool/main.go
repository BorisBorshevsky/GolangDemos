package main

import (
	"sync"

)

func main() {

	p := sync.Pool{}


	p.Put("sss")
	p.Put("bbb")

}
