package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	x := 5
	zero(&x)
	fmt.Println(x)
	one(xPtr)
	fmt.Println(*xPtr)
}
