package main

import "fmt"

func sum(nums []int, c chan int) {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	c <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int)

	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
