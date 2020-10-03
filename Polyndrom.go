package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println("Hello, playground")
	i := 1311
	//fmt.Printf("The binary value is %b\n", i)
	stringOfInteger := strconv.Itoa(i)
	length := len(stringOfInteger)
	j := 0
	var result bool = true
	for j < length {
		if stringOfInteger[j] != stringOfInteger[length-j-1] {
			result = false
			break
		}
		j++
	}
	if result {
		fmt.Printf("The integer %d is polyndrom\n", i)
	} else {
		fmt.Printf("The integer %d is not polyndrom\n", i)
	}
}
