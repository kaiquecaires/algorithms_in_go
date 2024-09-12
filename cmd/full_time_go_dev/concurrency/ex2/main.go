package main

import "fmt"

func main() {
	resCh := make(chan int, 10)
	resCh <- 1
	resCh <- 2
	close(resCh)

	for {
		num, ok := <-resCh
		if !ok {
			break
		}
		fmt.Println(num)
	}

	fmt.Println("Done!")
}
