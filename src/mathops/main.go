/*
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
*/
package main

import (
	"fmt"
	"math"
)

// this is a comment

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum", intSum)
	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum", floatSum)
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now", floatSum)
	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}
