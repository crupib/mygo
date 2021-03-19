package main

import (
	"fmt"
)

func main() {

	arr := [3]int{1, 2, 3}
	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)
	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 42, 27)
	/* count := 100
	for i := 0; i < count; i++ {
	  slice2 = append(slice2, i)
	}
	*/
	s2 := slice2[1:]
	s3 := slice2[:2]
	s4 := slice2[1:2]
	fmt.Println(slice2)
	fmt.Println(s2, s3, s4)
}
