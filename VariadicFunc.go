// Go program to illustrate the
// concept of variadic function
package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func addAll(elements ...int) int {
	retVal := 0
	for i := 0; i < len(elements); i++ {
		retVal = retVal + elements[i]*elements[i]*elements[i]
	}
	return retVal

}

func main() {

	// pass a slice in variadic function
	element := []string{"geeks", "FOR", "geeks"}
	fmt.Println(joinstr(element...))
	fmt.Println(addAll(1, 2))

}
