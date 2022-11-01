package main

import (
	"fmt"
)

func f1() int {
	return f2()
}
func f2() (r int) {
	r = 1
	return
}
func f3() (int, int) {
	return 5, 6
}
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	xs := []float64{98, 93, 77, 82, 83}
	ys := []int{1, 2, 5}
	fmt.Println(average(xs))
	fmt.Println(f1())
	x, y := f3()
	fmt.Println(x, y)
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(ys...))
}
