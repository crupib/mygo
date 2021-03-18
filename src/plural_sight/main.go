package main

import (
	"fmt"
)

const (
	first  = iota
	second = iota
)

func main() {
	fmt.Println("hello world")
	fmt.Println(first)
	fmt.Println(second)
}
