package main

import (
	"io"

	"time"

	"math"

	"reflect"
	"unsafe"

	"github.com/k0kubun/pp"
)

func main() {

	pr, pw := io.Pipe()

	payload := make([]byte, math.MaxInt8)

	var a = math.MaxInt64
	_ = a
	pp.Println(reflect.TypeOf(a).String())
	b := unsafe.Sizeof(a)
	pp.Println(int(b), b)

	go func() {
		defer pw.Close()
		pp.Println(pw.Write([]byte("Some data")))
	}()

	go func() {
		defer pr.Close()
		pp.Println(pr.Read(payload))
	}()

	time.Sleep(time.Second)
	pp.Println(string(payload))

}
