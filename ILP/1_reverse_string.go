package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func revWord(s string) string {
	sl := strings.Split(s, " ")
	var rev []string

	for i := len(sl) - 1; i >= 0; i-- {
		rev = append(rev, sl[i])

	}
	revStr := strings.Join(rev, " ")
	return revStr
}

func revSentence(s string) string {
	var rev string
	for i := len(s) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	return rev
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string you want to reverse")
	scanner.Scan()
	inputString := scanner.Text()
	fmt.Println("press 1 to reverse the words press 2 to reverse the whole string")
	scanner.Scan()
	option, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if option == 1 {
		result := revWord(inputString)
		fmt.Println("Reversed words: ", result)
	}
	if option == 2 {
		result := revSentence(inputString)
		fmt.Println("Reversed Sentence: ", result)
	}

}
