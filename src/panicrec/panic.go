package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
		fmt.Println("in func")
	}()
	panic("Panic")
	//	str := recover()
	//	fmt.Print(str)
}
