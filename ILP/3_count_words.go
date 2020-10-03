package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(s string) int {
	sl := strings.Split(s, " ")
	return len(sl)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string to be counted")
	scanner.Scan()
	input := scanner.Text()
	numWord := countWords(input)
	fmt.Println("There are", numWord, "Words in this sentence")
}
