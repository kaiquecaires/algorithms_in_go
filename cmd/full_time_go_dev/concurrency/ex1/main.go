package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	resCh := make(chan string, 10)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go FetchResourse(wg, resCh, i)
	}

	wg.Wait()
	close(resCh)

	for data := range resCh {
		fmt.Println(data)
	}
}

func FetchResourse(wg *sync.WaitGroup, resCh chan string, n int) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	resCh <- fmt.Sprintf("Number: %d", n)
}
