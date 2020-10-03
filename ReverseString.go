// Golang program to reverse a string
package main

// importing fmt
import (
	"fmt"
	"strconv"
)

// function, which takes a string as
// argument and return the reverse of string.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func main() {

	// Reversing the string.
	//str := "idappadi"
	num := 1122111
	str1 := strconv.Itoa(num)

	// returns the reversed string.
	strRev := reverse(str1)
	if str1 == strRev {
		fmt.Println(num, "is polyndrome")
	} else {
		fmt.Println(num, "is not polyndrome")
	}
}
