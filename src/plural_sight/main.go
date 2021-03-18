package main

import (
	"fmt"
)

const (
	first = iota
	second
)
const (
	third = iota
	fourth
)

func main() {
	fmt.Println("hello world")
	fmt.Println(first, second, third, fourth)

}
