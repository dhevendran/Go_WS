package main

import (
	"fmt"
	"sync"
)

var mem sync.Mutex
var wg sync.WaitGroup

func main() {
	for i := 'A'; i <= 'H'; i++ {
		mem.Lock()
		wg.Add(1)
		go func(i rune) {
			fmt.Println(string(i))
			mem.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
