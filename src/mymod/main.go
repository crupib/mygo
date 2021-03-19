package main

import (
	"fmt"
)

func main() {

	const pi = 3.1415
	const cc = 3
	const ccc int = 3
	var i int
	i = 42
	fmt.Println(i)
	var f float32 = 3.14
	fmt.Println(f)
	//	firstname := "Shit face"
	//	fmt.Println(firstname)
	b := true
	fmt.Println(b)
	c := complex(3, 4)
	fmt.Println(c)
	r, im := real(c), imag(c)
	fmt.Println(r, im)
	var firstname *string = new(string)
	*firstname = "bitch"
	fmt.Println(*firstname)
	secondname := "shit"
	fmt.Println(secondname)
	ptr := &secondname
	fmt.Println(ptr, *ptr)
	fmt.Println(pi)
	fmt.Println(cc + 3)
	fmt.Println(cc + 1.2)
	fmt.Println(float32(ccc) + 1.2)
}
