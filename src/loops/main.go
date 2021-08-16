package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops statement")
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	fmt.Println("for len")
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	fmt.Println("for range")
	for i := range colors {
		fmt.Println(colors[i])
	}
	fmt.Println("for no index")
	for _, color := range colors {
		fmt.Println(color)
	}
	fmt.Println("value")
	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto theEnds
		}
	}
theEnds:
	fmt.Println("End of program")

}
