package main

import "github.com/k0kubun/pp"

func main() {
	var a [2]int
	a[0] = 1
	a[1] = 1

	b := a[:]

	pp.Println("A", cap(a), len(a))
	pp.Println("B", cap(b), len(b))

	b = append(b, 3)

	pp.Println("A", cap(a), len(a))
	pp.Println("B", cap(b), len(b))
}
