package main

import (
	"fmt"
)

// Function to find average of cubes
func findAverageOfCube(n int) float64 {
	// Store sum of cubes of
	// numbers in the sum
	var sum float64 = 0

	// Calculate sum of cubes
	for i := 1; i <= n; i++ {
		sum += float64(i * i * i)
	}

	// Return average
	return sum / float64(n)
}

// Driver Code
func main() {
	for j := 0; j <= 10; j++ {
		avg := findAverageOfCube(j)
		fmt.Println(j, "---->", avg)
	}

}
