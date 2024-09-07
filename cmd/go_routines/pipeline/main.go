package main

import "fmt"

func generateNums(n int, c chan int) {
	for i := range n {
		c <- i
	}
	close(c)
}

func calcSquare(n chan int, c chan int) {
	for i := range n {
		c <- i * i
	}
	close(c)
}

func main() {
	generateNumsChan := make(chan int)
	calcSquareChan := make(chan int)

	go generateNums(10, generateNumsChan)
	go calcSquare(generateNumsChan, calcSquareChan)

	for i := range calcSquareChan {
		fmt.Println(i)
	}
}
