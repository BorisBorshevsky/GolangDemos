package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

var handler = func(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println("handler started")
	defer log.Println("handler ended")

	select {
	case <-time.After(2 * time.Second):
		fmt.Fprint(rw, "hello")

	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}
