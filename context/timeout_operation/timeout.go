package main

import (
	"time"

	"context"

	"github.com/k0kubun/pp"
)

func innerWork(ctx context.Context) (value, error) {
	time.Sleep(time.Millisecond * 500)
	pp.Println("working on", ctx.Value("key"))
	return ctx.Value("key"), nil

}

type value interface{}

func work(ctx context.Context, out chan<- value) error {
	v, err := innerWork(context.WithValue(ctx, "key", "Cool stuff for fetching data"))
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:

	}

	v, err = innerWork(context.WithValue(ctx, "key", "Cooler stuff for fetching data"))
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	v, err = innerWork(context.WithValue(ctx, "key", "Calculate and aggregate!! Yay"))
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case out <- v:
	}

	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	pp.Println("Hey, I'm going to do some work")

	result := make(chan value, 1)

	go work(ctx, result)

	select {
	case res := <-result:
		pp.Println("result is: ", res)
	case <-ctx.Done():
		pp.Println("Done due to: ", ctx.Err())
	}

	pp.Println("Finished. I'm going home")
}
