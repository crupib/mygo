package main

import "fmt"

// Calculates the sum of all the numbers in a slice and passes to the printSum function
func sumNumbers(numbers[]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	numbers := []int{5, 3, 4, 2, 5, 2, 62}
	sum := sumNumbers(numbers)
	fmt.Println(sum)
}

