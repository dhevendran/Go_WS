package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c3 <- true
	}()

	for i := 0; i < 1; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
		}
	}
}
