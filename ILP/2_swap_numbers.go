package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter X : ")
	scanner.Scan()
	x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println("Enter Y : ")
	scanner.Scan()
	y, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println("Before Swapping :", "X = ", x, "and Y = ", y)
	x = x + y
	y = x - y
	x = x - y
	fmt.Println("After Swapping :", "X = ", x, "and Y = ", y)
}
