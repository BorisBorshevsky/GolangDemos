package main

import (
	"sync"

	"github.com/k0kubun/pp"
)

func main() {

	p := sync.Pool{}


	p.Put("sss")
	p.Put("bbb")

	pp.Println(p.Get())
	pp.Println(p.Get())
	pp.Println(p.Get())

}
