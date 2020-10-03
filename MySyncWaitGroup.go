package main

import (
	"fmt"
	"strconv"
	"sync"
)

var w sync.WaitGroup

func myNewFun(str string) {
	fmt.Println("The string is ", str)
	w.Done()
}

func main() {
	str1, str2 := "Hello", "World"

	for i := 0; i < 10; i++ {
		w.Add(2)
		go myNewFun(str1 + " :" + strconv.Itoa(i))
		go myNewFun(str2 + " :" + strconv.Itoa(i))
	}

	w.Wait()
}
