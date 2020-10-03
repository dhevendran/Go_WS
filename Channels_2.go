package main

import "fmt"

func main() {
	c := make(chan int)
	go addAll(c, 1, 2, 5, 6)
	go addAll(c, 6, 7, 8, 10)
	sum1, sum2 := <-c, <-c
	fmt.Println("Sum1 ", sum1)
	fmt.Println("Sum2 ", sum2)
}

func addAll(ch chan int, ints ...int) {
	sum, v := 0, 0
	for _, v = range ints {
		sum = sum + v
	}
	ch <- sum
}
