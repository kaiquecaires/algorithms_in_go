package main

import (
	"fmt"
	"sync"
)

func goThrough(n int, c chan int) {
	for i := range n {
		c <- i
	}
	close(c)
}

func main() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()

	c := make(chan int)
	go goThrough(10, c)
	for i := range c {
		fmt.Println(i)
	}
}
