package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan rune, 50)
	for i := 'A'; i <= 'G'; i++ {
		wg.Add(1)
		ch <- i
		go printChar(ch)	
	}
	wg.Wait()
}

func printChar(ch chan rune) {
	fmt.Printf("%c ", <-ch)
	wg.Done()
}
