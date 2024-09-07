package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, sem chan bool) {
	defer wg.Done()
	sem <- true
	time.Sleep(2 * time.Second)
	fmt.Println(id)
	<-sem
}

func main() {
	wg := &sync.WaitGroup{}
	sem := make(chan bool, 2)
	for i := range 10 {
		wg.Add(1)
		go worker(i, wg, sem)
	}
	wg.Wait()
}
