package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mem sync.Mutex

func main() {
	ch := make(chan rune, 50)
	for i := 'A'; i <= 'G'; i++ {
		wg.Add(1)
		mem.Lock()
		ch <- i
		go printChar(ch)
	}
	wg.Wait()
}

func printChar(ch chan rune) {
	fmt.Printf("%c ", <-ch)
	mem.Unlock()
	wg.Done()
}
