package main

import (
	"time"

	"github.com/BorisBorshevsky/GolangDemos/catapult/api"
	"github.com/k0kubun/pp"
)

func main() {

	for i := 0; i < 3; i++ {

		start := time.Now()
		val, err := locationSvc.Alive()
		if err == nil {
			pp.Println("stop1", time.Since(start).Nanoseconds()/1e6, val.Commit)
		} else {
			pp.Println(err.Error())
		}
	}

	for i := 0; i < 0; i++ {

		start := time.Now()
		_, err := locationSvc.Alive()
		//if err != nil {
		//	pp.Println(err.Error(), val)
		//} else {
		//	pp.Println(val.Commit)
		//}

		if err == nil {
			pp.Println("stop1", time.Since(start).Nanoseconds()/1e6, err == nil)
		} else {
			pp.Println(err.Error())
		}
	}

}
