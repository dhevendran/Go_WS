package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 12344321
	a := strconv.Itoa(num)

	isPolyndrom := checkPolyndrom(a)
	if isPolyndrom {
		fmt.Println("The number =", num, " is polyndrom")
	} else {
		fmt.Println("The number =", num, " is NOT polyndrom")
	}

}

func checkPolyndrom(a string) bool {
	retValue := true
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		if a[i] != a[j] {
			retValue = false
			break
		}
	}
	return retValue
}
