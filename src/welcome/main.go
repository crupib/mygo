// Print a friendly greeting

package main

import (
	"fmt"
//	"unsafe"
)

func main() {
	fmt.Println("Welcome Gophers ☺")
/*    
	x, y := 1.0, 3.0
	
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)
	mean := (x+y)/2.0
	fmt.Printf("Result: %v, type of %T, %d\n",mean,mean,unsafe.Sizeof(mean))
	z := 10
	if z > 5 {
		fmt.Println("z is big")
	}
	if z > 100 {
		fmt.Println("z is very big")
	} else {
		fmt.Println("z is not that big")
	}
	if z > 5 && z < 15 {
		fmt.Println("z is just right")
	}
	if z < 20 || z > 30 {
		fmt.Println("z is out of range")
	}

	a := 11.0
	b := 20.0
	if frac := a/b; frac > 0.5 {
		fmt.Println("a is more that have of b")
	}

	x := 2
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}
	
	//fizz buzz
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println("fizz ",i)
		}
		if i%5 == 0 {
			fmt.Println("buzz", i)	
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz", i)	
		}
	}
	*/
	max := 0
	nums := []int{16,8,42,4,23,15}
	for _ , v:= range(nums) {
		//fmt.Printf("%v , %v\n",i,v)
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

}
