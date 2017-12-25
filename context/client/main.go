package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081", nil)
	//res, err := http.Get("http://localhost:8081")
	if err != nil {
		log.Fatal(err.Error())
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	io.Copy(os.Stdout, res.Body)
}
