package main

import "fmt"

// this is a comment

func main() {
	const myConst string = "Eat me!"
	var x = "Hello, Fucking World"
	fmt.Println(x)
	fmt.Println(myConst)
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
