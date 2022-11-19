package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
	sum := addValues(10, 8)
	fmt.Println("The sum is ", sum)
	multiSum, multiCount := addAllValues(4, 7, 9, 45)
	fmt.Println("Sum of multiple values:", multiSum)
	fmt.Println("Count of items", multiCount)
	subsum := subtractvalues(10, 5)
	fmt.Println("Sum of the subtraction is ", subsum)
}
func doSomething() {
	fmt.Println("Doing something")
}
func addValues(value1, value2 int) int {
	return value1 + value2
}
func subtractvalues(value1, value2 int) int {
	return value1 - value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
