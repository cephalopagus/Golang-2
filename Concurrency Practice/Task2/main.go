package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chanForResp := make(chan resp)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	go RPCCall(ctx, chanForResp)

	respo := <-chanForResp
	fmt.Println(respo.id, respo.err)

}

type resp struct {
	id  int
	err error
}

func RPCCall(ctx context.Context, ch chan<- resp) {

	select {
	case <-ctx.Done():
		ch <- resp{
			id:  0,
			err: errors.New("out of 4 second"),
		}

	case <-time.After(time.Duration(rand.Intn(10)) * time.Second):
		ch <- resp{
			id: rand.Intn(1000),
		}
	}

}
