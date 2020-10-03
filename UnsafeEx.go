package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct{}{}
	i := 7
	println(unsafe.Sizeof(a))
	fmt.Printf("Binary value of %d is %b", i, i)

}
