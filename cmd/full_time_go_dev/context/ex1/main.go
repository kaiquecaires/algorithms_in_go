package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	data, err := fetchUserId()
	fmt.Println(data, err)
}

func fetchUserId() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	type Response struct {
		data string
		err  error
	}

	responseCh := make(chan Response, 1)

	go func() {
		res, err := thirdPartyHttpCall()
		responseCh <- Response{
			data: res,
			err:  err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case response := <-responseCh:
		return response.data, response.err
	}
}

func thirdPartyHttpCall() (string, error) {
	time.Sleep(1 * time.Second)
	return "some data", nil
}
