package main

import (
	"fmt"
	"math"
)

func main() {
	var anInt int = 5
	var aFloat float64 = 42
	sum := float64(anInt) + aFloat
	fmt.Printf("Sum: %v, Type: %T\n", sum, sum)

	var aFloat2 = 3.14259
	var rounded = math.Round(aFloat2)
	fmt.Printf("Original: %v, Rounded: %v\n", aFloat2, rounded)

	var aFloat3 = math.Pi
	var rounded2 = math.Round(aFloat3)
	fmt.Printf("Original: %v, Rounded: %v\n", aFloat3, rounded2)
}
